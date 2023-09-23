/** @type {import('tailwindcss').Config}*/
const config = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      fontFamily: {
        'sans-pro': ['Source Sans Pro', 'sans-serif'],
        'open-sans': ['Open Sans Variable', 'sans-serif'],
      },
    },
  },
  darkMode: 'class',
  plugins: [require('@tailwindcss/typography')],
};

module.exports = config;
