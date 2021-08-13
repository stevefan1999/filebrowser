<template>
  <div>
    <header-bar v-if="error || req.type == null" showLogo showMenu />

    <b-row align-v="center">
      <b-col sm="12" md="8" lg="9" xl="10">
        <breadcrumbs base="/files" />
      </b-col>
      <b-col sm="12" md="4" lg="3" xl="2">
        <b-row align-v="center">
          <b-col md="4" lg="4"> Disk Volume:</b-col>
          <b-col md="8" lg="8">
            <b-progress
              v-if="req.disk_stat"
              :max="req.disk_stat.total"
              :variant="diskBarVariant"
              height="2rem"
            >
              <b-progress-bar
                class="overflow-visible"
                :value="req.disk_stat.used"
              >
                <span style="color: black">
                  {{ prettyBytes(req.disk_stat.used, { binary: true }) }} /
                  {{ prettyBytes(req.disk_stat.total, { binary: true }) }}
                </span>
              </b-progress-bar>
            </b-progress>
          </b-col>
        </b-row>
      </b-col>
    </b-row>

    <errors v-if="error" :errorCode="error.message" />
    <component :is="currentView" v-else-if="currentView"></component>
    <div v-else>
      <h2 class="message delayed">
        <div class="spinner">
          <div class="bounce1"></div>
          <div class="bounce2"></div>
          <div class="bounce3"></div>
        </div>
        <span>{{ $t("files.loading") }}</span>
      </h2>
    </div>
  </div>
</template>

<script>
import { files as api } from "@/api";
import { mapMutations, mapState } from "vuex";

import HeaderBar from "@/components/header/HeaderBar";
import Breadcrumbs from "@/components/Breadcrumbs";
import Errors from "@/views/Errors";
import Preview from "@/views/files/Preview";
import Listing from "@/views/files/Listing";
import Editor from "@/views/files/Editor";
import prettyBytes from "pretty-bytes";

function clean(path) {
  return path.endsWith("/") ? path.slice(0, -1) : path;
}

export default {
  name: "files",
  components: {
    HeaderBar,
    Breadcrumbs,
    Errors,
    Preview,
    Listing,
    Editor,
  },
  data: function () {
    return {
      error: null,
      width: window.innerWidth,
      progress: 25,
      refreshTimerHandle: null
    };
  },
  computed: {
    ...mapState(["req", "reload", "loading", "show"]),
    currentView() {
      if (this.req.type == undefined) {
        return null;
      }

      if (this.req.isDir) {
        return "listing";
      } else if (
        this.req.type === "text" ||
        this.req.type === "textImmutable"
      ) {
        return "editor";
      } else {
        return "preview";
      }
    },
    diskUsedPercentage() {
      return this.req.disk_stat
        ? (this.req.disk_stat.used / this.req.disk_stat.total).toFixed(2)
        : 0;
    },
    diskBarVariant() {
      if (this.diskUsedPercentage >= 0.75) {
        return "danger";
      }
      if (this.diskUsedPercentage >= 0.5) {
        return "warning";
      }
      return "info";
    },
  },
  created() {
    this.fetchData();
  },
  watch: {
    $route: "fetchData",
    reload: function (value) {
      if (value === true) {
        this.fetchData();
      }
    },
  },
  mounted() {
    window.addEventListener("keydown", this.keyEvent);
    this.refreshTimerHandle = setInterval(async () => {
      await this.updateData();
    }, 5000);
  },
  beforeDestroy() {
    window.removeEventListener("keydown", this.keyEvent);
    if (this.refreshTimerHandle !== null) {
      clearInterval(this.refreshTimerHandle);
      this.refreshTimerHandle = null;
    }
  },
  destroyed() {
    if (this.$store.state.showShell) {
      this.$store.commit("toggleShell");
    }
    this.$store.commit("updateRequest", {});
  },
  methods: {
    ...mapMutations(["setLoading"]),
    prettyBytes,
    async updateData() {
      let url = this.$route.path;
      if (url === "") url = "/";
      if (url[0] !== "/") url = "/" + url;

      try {
        const res = await api.fetch(url);

        if (clean(res.path) !== clean(`/${this.$route.params.pathMatch}`)) {
          return;
        }

        this.$store.commit("updateRequest", res);
      } catch (e) {
        // do nothing
      }
    },
    async fetchData() {
      // Reset view information.
      this.$store.commit("setReload", false);
      this.$store.commit("resetSelected");
      this.$store.commit("multiple", false);
      this.$store.commit("closeHovers");

      // Set loading to true and reset the error.
      this.setLoading(true);
      this.error = null;

      let url = this.$route.path;
      if (url === "") url = "/";
      if (url[0] !== "/") url = "/" + url;

      try {
        const res = await api.fetch(url);

        if (clean(res.path) !== clean(`/${this.$route.params.pathMatch}`)) {
          return;
        }

        this.$store.commit("updateRequest", res);
        document.title = `${res.name} - ${this.$route.name}`;
      } catch (e) {
        this.error = e;
      } finally {
        this.setLoading(false);
      }
    },
    keyEvent(event) {
      // F1!
      if (event.keyCode === 112) {
        event.preventDefault();
        this.$store.commit("showHover", "help");
      }
    },
  },
};
</script>
<style>
.capacity {
  border-bottom: 1px none;
  /*display: flex;*/
  align-items: center;
}
</style>
