var path = require('path');
var merge = require('webpack-merge');
var bower_components = path.resolve(__dirname, 'bower_components');

var config = {
  entry: './js/app.js',
  output: {
    path: 'dist',
    filename: 'bundle.js'
  },
  module: {
    loaders: [{
      test: /\.jsx?$/,
      exclude: /node_modules/,
      loader: 'babel-loader',
      query: {
        presets: ['react', 'es2015']
      }
    }, {
      test: /\.scss$/,
      exclude: /node_modules/,
      loaders: ["style", "css", "sass"]
    }, {
      test: /\.css$/,
      exclude: /node_modules/,
      loaders: ["style", "css"]
    }, {
      test: /\.png$/,
      exclude: /node_modules/,
      loader: "file-loader"
    }, {
      test: /\.(woff|svg|ttf|eot)([\?]?.*)$/,
      exclude: /node_modules/,
      loader: 'file-loader?name=[name].[ext]'
    }]
  },
  resolve: {
    alias: {
      'bootstrap': bower_components + '/bootstrap/dist/css/bootstrap.css'
    }
  }
};

module.exports = config;
