const zlib = require("zlib");
const CompressionPlugin = require("compression-webpack-plugin");

const algorithms = {
  brotli: {
    filename: "[file].br[query]",
    algorithm: "brotliCompress",
    include: /\.(js|css|html|svg|json|ttf|eot|woff2?)(\?.*)?$/i,
    compressionOptions: {
      params: {
        [zlib.constants.BROTLI_PARAM_QUALITY]: 11,
      },
    },
    minRatio: 0.8,
  },
  gzip: {
    filename: "[file].gz[query]",
    algorithm: "gzip",
    include: /\.(js|css|html|svg|json|ttf|eot|woff2?)(\?.*)?$/i,
    minRatio: 0.8,
  },
};

module.exports = {
  runtimeCompiler: true,
  publicPath: "[{[ .StaticURL ]}]",

  chainWebpack(config) {
    for (const [name, algorithm] of Object.entries(algorithms)) {
      config
        .plugin(`${name}-compression`)
        .use(CompressionPlugin)
        .init((Plugin) => new Plugin(algorithm));
    }
  },
};
