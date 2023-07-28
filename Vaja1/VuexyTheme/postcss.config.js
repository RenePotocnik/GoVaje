const rtl = require('postcss-rtl')

module.exports = {
  plugins: [
      rtl(),
      require('tailwindcss')('./tailwind.config.js'),
      require('autoprefixer')
    ]
}
