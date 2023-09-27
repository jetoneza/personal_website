import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { HTTP_CODE_NOT_FOUND } from '$lib/constants';

export const load: PageServerLoad = async ({ fetch, params }) => {
  const { slug } = params;

  // TODO: Add error handling
  const res = await fetch(`/api/v1/posts/slug/${slug}`);
  const jsonResponse = await res.json();
  const post = jsonResponse.data;

  if (!post) {
    throw error(HTTP_CODE_NOT_FOUND, { message: 'Page not found.' });
  }

  return {
    post,
  };
};
