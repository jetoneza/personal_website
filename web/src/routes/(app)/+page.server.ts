import { formatDate } from '$lib/utils/date';
import type { Actions, PageServerLoad } from './$types';

const TIME_YEAR = 60 * 60 * 24 * 365;

export const load: PageServerLoad = async ({ fetch }) => {
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

  // TODO: Add image
  const meta = {
    title: 'Jet Ordaneza',
    description: `A Software Engineer. I love building cool things with code.
      I'm always curious about how things work and love exploring new stuff.`,
    keywords: 'software engineer,reactjs,react,nodejs,node,rust,go,golang,tailwindcss',
    url: 'https://www.jetordaneza.com',
    imageUrl: '',
  }

  return { posts, meta };
};

export const actions: Actions = {
  setTheme: async ({ url, cookies }) => {
    const theme = url.searchParams.get('theme');

    if (theme) {
      cookies.set('theme', theme, {
        path: '/',
        httpOnly: true,
        maxAge: TIME_YEAR,
        secure: false, // TODO: Set to true if ssl is setup
      });
    }
  },
};
