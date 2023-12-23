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

    const publishedAt = new Date(data.get('published_at') as string);

    const response = await fetch('/api/v1/posts', {
      method: 'POST',
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
      return fail(HTTP_CODE_BAD_REQUEST, {
        ...result,
        message: MESSAGE_POST_CREATION_ERROR,
      });
    }

    redirect(
      HTTP_CODE_SEE_OTHER,
      `/admin/dashboard/blog?message={"type":"success","value":"${MESSAGE_POST_CREATION_SUCCESS}"}`
    );
  },
};
