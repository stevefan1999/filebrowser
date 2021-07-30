package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/wesovilabs/koazee"

	"github.com/filebrowser/filebrowser/v2/auth"
	"github.com/filebrowser/filebrowser/v2/settings"
	"github.com/filebrowser/filebrowser/v2/storage"
	"github.com/filebrowser/filebrowser/v2/version"
)

func handleWithStaticData(w http.ResponseWriter, _ *http.Request, d *data, fSys fs.FS, file, contentType string) (int, error) {
	w.Header().Set("Content-Type", contentType)

	auther, err := d.store.Auth.Get(d.settings.AuthMethod)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	data := map[string]interface{}{
		"Name":            d.settings.Branding.Name,
		"DisableExternal": d.settings.Branding.DisableExternal,
		"BaseURL":         d.server.BaseURL,
		"Version":         version.Version,
		"StaticURL":       path.Join(d.server.BaseURL, "/static"),
		"Signup":          d.settings.Signup,
		"NoAuth":          d.settings.AuthMethod == auth.MethodNoAuth,
		"AuthMethod":      d.settings.AuthMethod,
		"LoginPage":       auther.LoginPage(),
		"CSS":             false,
		"ReCaptcha":       false,
		"Theme":           d.settings.Branding.Theme,
		"EnableThumbs":    d.server.EnableThumbnails,
		"ResizePreview":   d.server.ResizePreview,
		"EnableExec":      d.server.EnableExec,
	}

	if d.settings.Branding.Files != "" {
		fPath := filepath.Join(d.settings.Branding.Files, "custom.css")
		_, err := os.Stat(fPath) //nolint:govet

		if err != nil && !os.IsNotExist(err) {
			log.Printf("couldn't load custom styles: %v", err)
		}

		if err == nil {
			data["CSS"] = true
		}
	}

	if d.settings.AuthMethod == auth.MethodJSONAuth {
		raw, err := d.store.Auth.Get(d.settings.AuthMethod) //nolint:govet
		if err != nil {
			return http.StatusInternalServerError, err
		}

		auther := raw.(*auth.JSONAuth)

		if auther.ReCaptcha != nil {
			data["ReCaptcha"] = auther.ReCaptcha.Key != "" && auther.ReCaptcha.Secret != ""
			data["ReCaptchaHost"] = auther.ReCaptcha.Host
			data["ReCaptchaKey"] = auther.ReCaptcha.Key
		}
	}

	b, err := json.Marshal(data)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	data["Json"] = string(b)

	fileContents, err := fs.ReadFile(fSys, file)
	if err != nil {
		if err == os.ErrNotExist {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}
	index := template.Must(template.New("index").Delims("[{[", "]}]").Parse(string(fileContents)))
	err = index.Execute(w, data)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func getStaticHandlers(store *storage.Storage, server *settings.Server, assetsFs fs.FS) (index, static http.Handler) {
	index = handle(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
		if r.Method != http.MethodGet {
			return http.StatusNotFound, nil
		}

		w.Header().Set("x-xss-protection", "1; mode=block")
		return handleWithStaticData(w, r, d, assetsFs, "index.html", "text/html; charset=utf-8")
	}, "", store, server)

	static = handle(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
		if r.Method != http.MethodGet {
			return http.StatusNotFound, nil
		}

		const maxAge = 86400 // 1 day
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%v", maxAge))

		if d.settings.Branding.Files != "" {
			if strings.HasPrefix(r.URL.Path, "img/") {
				fPath := filepath.Join(d.settings.Branding.Files, r.URL.Path)
				if _, err := os.Stat(fPath); err == nil {
					http.ServeFile(w, r, fPath)
					return 0, nil
				}
			} else if r.URL.Path == "custom.css" && d.settings.Branding.Files != "" {
				http.ServeFile(w, r, filepath.Join(d.settings.Branding.Files, "custom.css"))
				return 0, nil
			}
		}

		if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		}

		found, errCode, err := tryCompressedFile(".br", "br", assetsFs, r, w)
		if errCode != 0 || err != nil {
			return errCode, err
		}
		if found {
			return 0, nil
		}

		found, errCode, err = tryCompressedFile(".gz", "gzip", assetsFs, r, w)
		if errCode != 0 || err != nil {
			return errCode, err
		}
		if found {
			return 0, nil
		}

		http.FileServer(http.FS(assetsFs)).ServeHTTP(w, r)
		return 0, nil
	}, "/static/", store, server)

	return index, static
}

func tryCompressedFile(
	suffix, encodingType string, assetsFs fs.FS, r *http.Request, w http.ResponseWriter,
) (found bool, statusCode int, err error) {
	acceptedEncodings := strings.Split(r.Header["Accept-Encoding"][0], ",")
	acceptEncodings := koazee.StreamOf(acceptedEncodings).Map(strings.TrimSpace).Do()
	acceptedEncoding, err := acceptEncodings.Contains(encodingType)
	err = errors.Unwrap(err)
	if err != nil {
		return false, 0, err
	}
	if !acceptedEncoding {
		return false, 0, nil
	}
	fileContents, err := fs.ReadFile(assetsFs, r.URL.Path+suffix)
	if err != nil {
		return false, 0, err
	}
	w.Header().Set("Content-Encoding", encodingType)
	_, err = w.Write(fileContents)
	if err != nil {
		return true, http.StatusInternalServerError, err
	}
	return true, 0, nil
}
