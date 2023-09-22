/** @type {import('tailwindcss').Config}*/
const config = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {},
  },
  darkMode: 'class',
  plugins: [require('@tailwindcss/typography')],
};

module.exports = config;
