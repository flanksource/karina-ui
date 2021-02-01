module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  "devServer": {
    "disableHostCheck": true,
    port: 8085,
    proxy: {
      '^/api': {
        target: 'http://localhost:8080/',
        changeOrigin: true,
        logLevel: 'debug'
      },
    }
  }
}