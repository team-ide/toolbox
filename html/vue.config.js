module.exports = {
  assetsDir: process.env.assetsDir,
  publicPath: process.env.publicPath,
  "devServer": {
    "port": 8000,
    "proxy": {
      "/server": {
        target: "http://127.0.0.1:18000/toolbox",
        changeOrigin: true
      },
    }
  },
  "productionSourceMap": false,
  "transpileDependencies": [
  ],
}