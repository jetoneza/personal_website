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

    try {
      const response = await fetch(`/api/v1/posts/${id}`, {
        method: 'PUT',
        headers: {
          cookie: request.headers.get('cookie') as string,
        },
        body: data,
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
