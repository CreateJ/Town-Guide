const {resolve} = require('path')
const CopyWebpackPlugin = require('copy-webpack-plugin')
const {CleanWebpackPlugin} = require('clean-webpack-plugin')

module.exports = {
  context: resolve('src'),
  entry: {
    'app': './app.js',
    'pages/index/index': './pages/index/index.js',
    'pages/map/index': './pages/map/index.js',
    'pages/visits/index': './pages/visits/index.js',
    'pages/mine/index': './pages/mine/index.js',
    'pages/logs/logs': './pages/logs/logs.js'
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        use: 'babel-loader'
      }
    ]
  },
  output: {
    path: resolve('dist'),
    filename: '[name].js',
  },
  plugins: [
    new CleanWebpackPlugin({
      cleanStaleWebpackAssets: false,
    }),
    new CopyWebpackPlugin({
      patterns: [
        {
          from: '**/*',
          to: './',
          filter(fileName) {
            return fileName.substring(fileName.length - 3) !== '.js'
          },
        },
      ],
    }),
  ],
  mode: 'none',
}