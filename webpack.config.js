var path = require('path');
var webpack = require('webpack');

module.exports = {
  // plugins:[
  //   new webpack.DefinePlugin({
  //     'process.env':{
  //       'NODE_ENV': JSON.stringify('production')
  //     }
  //   }),
  //   new webpack.optimize.UglifyJsPlugin({
  //     compress:{
  //       warnings: true
  //     }
  //   })
  // ],
  entry: './react/src/main.jsx',
  output: { path: __dirname, filename: 'public/bundle.js' },
  module: {
    loaders: [
      {
        test: /.jsx?$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
        query: {
          presets: ['es2015', 'react']
        }
      }
    ]
  },
};