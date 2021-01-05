const { useGuarkLockFile, checkBeforeBuild } = require('guark/build')
const VueLoaderPlugin = require('vue-loader/lib/plugin')

checkBeforeBuild()

module.exports =
    {
        outputDir: process.env.GUARK_BUILD_DIR,
        productionSourceMap: process.env.NODE_ENV === 'production' ? false : true,
        configureWebpack:
          {
            devServer:
                {
                  // After server started you should call useGuarkLockFile.
                  after: (app, server, compiler) => compiler.hooks.done.tap("Guark", useGuarkLockFile)
                },
              mode: 'development',
              module: {
                  rules: [
                      { test: /\.vue$/, include: /src/, loader: 'vue-loader', options: { loaders: { js: 'awesome-typescript-loader?silent=true', loader: "unicode-loader" } } },
                      // { test: /\.ts$/, include: /ClientApp/, use: 'awesome-typescript-loader?silent=true' },
                      { test: /\.css$/, use: ['style-loader', 'css-loader'] },
                      { test: /\.(png|jpg|jpeg|gif|svg|woff2|woff)$/, include: /src/, use: 'url-loader?limit=25000' }
                  ]
              },
              plugins: [
                  // make sure to include the plugin for the magic
                  new VueLoaderPlugin()
              ]
          },
        transpileDependencies: [
          "vuetify"
        ]
    }
