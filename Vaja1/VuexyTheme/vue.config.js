const path = require('path')
const WebpackShellPluginNext = require('webpack-shell-plugin-next')

// if it's a production build, then remove icons,
// otherwise execute empty function
// eslint-disable-next-line multiline-ternary
const removeIconsOnProductionBuild = (process.env.NODE_ENV === 'production')
  ? new WebpackShellPluginNext({
    onBuildStart: {
      scripts: [
        'echo "--- DELETING icons FROM bootstrap-vue icons.js ---"',
        'bash scripts/remove-icons-from-bootstrap-vue.sh'
      ],
      blocking: true,
      parallel: false
    }
  }) : () => {}

module.exports = {
  publicPath: '/',
  lintOnSave: false,
  css: {
    loaderOptions: {
      sass: {
        sassOptions: {
          includePaths: ['./node_modules', './src/assets']
        }
      }
    }
  },
  configureWebpack: {
    resolve: {
      alias: {
        '@themeConfig': path.resolve(__dirname, 'themeConfig.js'),
        '@core': path.resolve(__dirname, 'src/@core'),
        '@validations': path.resolve(__dirname, 'src/@core/utils/validations/validations.js'),
        '@axios': path.resolve(__dirname, 'src/libs/axios')
      }
    },
    plugins: [
      // bootstrap-vue adds icons.js, which is 600kb after being bundled (for production), even though, icons aren't used anywhere.
      // on production build, regex command is executed on 'node_modules/bootstrap-vue/esm/icons/icons.js', where icons are defined,
      // which removes all icons expect few, that are used internally by bootstrap in calendar, time, avatar,...
      removeIconsOnProductionBuild
    ]
  },
  chainWebpack: config => {
    config.module
      .rule('vue')
      .use('vue-loader')
      .loader('vue-loader')
      .tap(options => {
        // eslint-disable-next-line no-param-reassign
        options.transformAssetUrls = {
          'img': 'src',
          'image': 'xlink:href',
          'b-avatar': 'src',
          'b-img': 'src',
          'b-img-lazy': ['src', 'blank-src'],
          'b-card': 'img-src',
          'b-card-img': 'src',
          'b-card-img-lazy': ['src', 'blank-src'],
          'b-carousel-slide': 'img-src',
          'b-embed': 'src'
        }
        return options
      })
  },
  transpileDependencies: ['resize-detector'],
  pluginOptions: {
    webpackBundleAnalyzer: {
      openAnalyzer: false
    }
  }

}
