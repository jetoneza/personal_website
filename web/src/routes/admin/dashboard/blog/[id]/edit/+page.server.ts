import { API_STATUS, HTTP_CODE_BAD_REQUEST, MESSAGE_POST_CREATION_ERROR } from '$lib/constants';
import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, params }) => {
  const { id } = params;

  // TODO: Add error handling
  const res = await fetch(`/api/v1/posts/${id}`);
  const jsonResponse = await res.json();
  const post = jsonResponse.data;

  return {
    post,
  };
};

export const actions: Actions = {
  default: async ({ request, fetch, params }) => {
    const data = await request.formData();

    const { id } = params

    try {
      const response = await fetch(`/api/v1/posts/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          title: data.get('title'),
          content: data.get('content'),
          slug: data.get('slug'),
          category: data.get('category'),
          description: data.get('description'),
          meta_title: data.get('meta_title'),
          meta_description: data.get('meta_description'),
          meta_keyword: data.get('meta_keyword'),
        }),
      });

      const result = JSON.parse(await response.text());

      if (result.status === API_STATUS.FAIL) {
        throw new Error(MESSAGE_POST_CREATION_ERROR)
      }
    } catch (_) {
      return fail(HTTP_CODE_BAD_REQUEST, {
        status: API_STATUS.FAIL,
        message: MESSAGE_POST_CREATION_ERROR,
      });
    }
  },
};
