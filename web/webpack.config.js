const path = require('path');

module.exports = {
  entry: './js/main.js',
  output: {
    filename: 'index.js',
    path: path.resolve(__dirname, 'js'),
  },
  mode: 'production'
};