const path = require('path')
const dotenv = require('dotenv')

dotenv.config() // setup .env

const IS_PROD = process.env.NODE_ENV === 'production'

function resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = {
  lintOnSave: false, // process.env.NODE_ENV !== 'production',
  pages: {
    app: {
      title: 'mydrive',
      entry: 'src/renderer/index.js',
      template: 'public/index.html',
      filename: 'index.html'
    }
  },
  publicPath: './',
  assetsDir: 'assets',
  outputDir: 'dist',
  productionSourceMap: !IS_PROD,
  devServer: {
    // can be overwritten by process.env.HOST
    host: 'localhost',
    port: 8099
  },
  chainWebpack: (config) => {
    // Path alias, such as "@" for "src", etc.
    config.resolve.alias
      .set('@', resolve('src'))
      .set('src', resolve('src'))
      .set('assets', resolve('src/assets'))
      .set('plugins', resolve('src/plugins'))
      .set('svg', resolve('src/assets/svg'))
      .set('locales', resolve('src/locales'))
      .set('backend', resolve('src/backend'))
      .set('renderer', resolve('src/renderer'))
      .set('main', resolve('src/main'))
      .set('views', resolve('src/renderer/views'))
      .set('components', resolve('src/renderer/components'))
    // svg config
    const svgRule = config.module.rule('svg')
    svgRule.uses.clear()
    svgRule
      .use('svg-sprite-loader')
      .loader('svg-sprite-loader')
      .options({
        symbolId: 'icon-[name]'
      })
  },
  pluginOptions: {
    electronBuilder: {
      mainProcessFile: 'src/main/index.js',
      mainProcessWatch: ['src/main'],
      nodeIntegration: true,
      builderOptions: {
        win: {
          icon: './public/favicon.ico'
        },
        productName: 'mydrive'
      }
    },
    // i18n config
    i18n: {
      locale: process.env.VUE_APP_DEFAULT_LANGUAGE,
      fallbackLocale: 'en',
      localeDir: 'locales',
      enableInSFC: false
    }
  }
}
