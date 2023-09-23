import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, params }) => {
  const { slug } = params;

  // TODO: Add error handling
  const res = await fetch(`/api/v1/posts/${slug}`);
  const jsonResponse = await res.json();
  const post = jsonResponse.data;

  return {
    post,
  };
};
