/** @type {import('tailwindcss').Config} */
module.exports = {
  // todo: add html and go file types
  content: [
    "./templates/**/*.templ",
    "./templates/*.templ",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('daisyui'),
  ],
}

