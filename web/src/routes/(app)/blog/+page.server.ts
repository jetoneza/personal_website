import { formatDate } from '$lib/utils/date';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, url }) => {
  const page = Number(url.searchParams.get('page') ?? '1');
  const limit = Number(url.searchParams.get('limit') ?? '10');

  const qs = new URLSearchParams();
  qs.set('page', page.toString());
  qs.set('limit', limit.toString());

  // TODO: Add error handling
  const res = await fetch(`/api/v1/posts/published?${qs.toString()}`);
  const jsonResponse = await res.json();

  // TODO: Improve formatting of dates
  const posts = jsonResponse.data.map((post: Post) => {
    return {
      ...post,
      formattedPublishedAt: formatDate((post.publishedAt || post.createdAt) as string),
    };
  });
  //
  // TODO: Add image
  const meta = {
    title: 'Blog | Jet Ordaneza',
    description: 'Check out my latest blog posts where we explore my thoughts about anything tech.',
    keywords: 'software engineer,reactjs,react,nodejs,node,rust,go,golang,tailwindcss',
    url: 'https://www.jetordaneza.com/blog',
    imageUrl: '',
  };

  return { posts, meta };
};
