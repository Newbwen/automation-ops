//const { defineConfig } = require('@vue/cli-service')
module.exports = {
  devServer: {
    port:3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',//跨域请求的地址
        changeOrigin: true,//是否跨域
        pathRewrite: {
          '^/api': ''//路径重写，将/api替换成空字符串
        }
      }
    }

},
chainWebpack(config) {
  config.plugin('html').tap(args => { // 修改htmlWebpackPlugin的title
    args[0].title = 'automation-ops-web'
    return args
  })
}
}