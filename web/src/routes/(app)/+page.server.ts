import type { Actions } from './$types';

const TIME_YEAR = 60 * 60 * 24 * 365;

export const actions: Actions = {
  setTheme: async ({ url, cookies }) => {
    const theme = url.searchParams.get('theme');

    if (theme) {
      cookies.set('theme', theme, {
        path: '/',
        maxAge: TIME_YEAR,
      });
    }
  },
};
