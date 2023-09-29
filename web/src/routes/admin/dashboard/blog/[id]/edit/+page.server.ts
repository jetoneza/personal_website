import {
  API_STATUS,
  HTTP_CODE_BAD_REQUEST,
  HTTP_CODE_SEE_OTHER,
  MESSAGE_POST_CREATION_ERROR,
  MESSAGE_POST_UPDATE_SUCCESS,
} from '$lib/constants';
import { fail, redirect } from '@sveltejs/kit';
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

    const { id } = params;

    const publishedAt = new Date(data.get('published_at') as string);

    try {
      const response = await fetch(`/api/v1/posts/${id}`, {
        method: 'PUT',
        headers: {
          cookie: request.headers.get('cookie') as string,
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
          meta_keywords: data.get('meta_keywords'),
          meta_image_url: data.get('meta_image_url'),
          published: data.get('published'),
          published_at: publishedAt,
        }),
      });

      const result = JSON.parse(await response.text());

      if (result.status === API_STATUS.FAIL) {
        throw new Error(MESSAGE_POST_CREATION_ERROR);
      }
    } catch (_) {
      return fail(HTTP_CODE_BAD_REQUEST, {
        status: API_STATUS.FAIL,
        message: MESSAGE_POST_CREATION_ERROR,
      });
    }

    throw redirect(
      HTTP_CODE_SEE_OTHER,
      `/admin/dashboard/blog?message={"type":"success","value":"${MESSAGE_POST_UPDATE_SUCCESS}"}`
    );
  },
};
