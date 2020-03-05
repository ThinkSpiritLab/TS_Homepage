module.exports = {
  publicPath: './',
  outputDir: 'dist',
  lintOnSave: true,
  productionSourceMap: false,
  devServer: {
    host: '127.0.0.1',
    port:'8501',
    disableHostCheck: true,
    https:false,
    hotOnly: false,
    //代理
    proxy: 'http://localhost:8500/api/v1/'
  }
}
