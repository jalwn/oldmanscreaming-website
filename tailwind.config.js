/** @type {import('tailwindcss').Config} */
module.exports = {
  // todo: add html and go file types
  content: [
    "./internal/**/*.{go,js,templ,html}"
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('daisyui'),
  ],
}

