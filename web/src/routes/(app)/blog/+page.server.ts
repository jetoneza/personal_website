import type { PageServerLoad } from './$types';

const monthNames = [
  'January',
  'February',
  'March',
  'April',
  'May',
  'June',
  'July',
  'August',
  'September',
  'October',
  'November',
  'December',
];

export const load: PageServerLoad = async ({ fetch, url }) => {
  const page = Number(url.searchParams.get('page') ?? '1');
  const limit = Number(url.searchParams.get('limit') ?? '10');

  const qs = new URLSearchParams();
  qs.set('page', page.toString());
  qs.set('limit', limit.toString());

  // TODO: Add error handling
  const res = await fetch(`/api/v1/posts?${qs.toString()}`);
  const jsonResponse = await res.json();

  // TODO: Improve formatting of dates
  const posts = jsonResponse.data.map((post: Post) => {
    const date = new Date(post.createdAt as string);
    return {
      ...post,
      formattedCreatedAt: `${monthNames[date.getMonth()]} ${date.getDay()}, ${date.getFullYear()}`,
    };
  });

  return { posts };
};
