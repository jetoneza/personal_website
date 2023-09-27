import { formatDate } from '$lib/utils/date';
import type { Actions, PageServerLoad } from './$types';

const TIME_YEAR = 60 * 60 * 24 * 365;

export const load: PageServerLoad = async ({ fetch, url }) => {
  const qs = new URLSearchParams();
  qs.set('page', '1');
  qs.set('limit', '3');

  // TODO: Add error handling
  const res = await fetch(`/api/v1/posts?${qs.toString()}`);
  const jsonResponse = await res.json();

  // TODO: Improve formatting of dates
  const posts = jsonResponse.data.map((post: Post) => {
    return {
      ...post,
      formattedCreatedAt: formatDate(post.createdAt as string),
    };
  });

  return { posts };
};

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
