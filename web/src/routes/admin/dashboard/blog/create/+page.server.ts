import { fail, redirect } from '@sveltejs/kit';

// Types
import type { Actions } from './$types';

// Constants
import {
  API_STATUS,
  HTTP_CODE_BAD_REQUEST,
  HTTP_CODE_SEE_OTHER,
  MESSAGE_POST_CREATION_ERROR,
  MESSAGE_POST_CREATION_SUCCESS,
} from '$lib/constants';

export const actions: Actions = {
  default: async ({ request, fetch }) => {
    const data = await request.formData();

    const response = await fetch('/api/v1/posts', {
      method: 'POST',
      headers: {
        cookie: request.headers.get('cookie') as string,
      },
      body: data,
    });

    const result = JSON.parse(await response.text());

    if (result.status === API_STATUS.FAIL) {
      return fail(HTTP_CODE_BAD_REQUEST, {
        ...result,
        message: MESSAGE_POST_CREATION_ERROR,
      });
    }

    throw redirect(
      HTTP_CODE_SEE_OTHER,
      `/admin/dashboard/blog?message={"type":"success","value":"${MESSAGE_POST_CREATION_SUCCESS}"}`
    );
  },
};
