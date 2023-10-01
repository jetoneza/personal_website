const disabledCss = {
  code: false,
  'code::before': false,
  'code::after': false,
};

/** @type {import('tailwindcss').Config}*/
const config = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      fontFamily: {
        'sans-pro': ['Source Sans Pro', 'sans-serif'],
        'open-sans': ['Open Sans Variable', 'sans-serif'],
      },
      typography: {
        DEFAULT: { css: disabledCss },
        sm: { css: disabledCss },
        lg: { css: disabledCss },
        xl: { css: disabledCss },
        '2xl': { css: disabledCss },
      },
    },
  },
  darkMode: 'class',
  plugins: [require('@tailwindcss/typography')],
};

module.exports = config;
