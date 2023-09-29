import { fail } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { API_STATUS, HTTP_CODE_BAD_REQUEST, MESSAGE_POST_FETCH_ERROR } from '$lib/constants';

export const load: PageServerLoad = async ({ request, fetch, url }) => {
  const page = Number(url.searchParams.get('page') ?? '1');
  const limit = Number(url.searchParams.get('limit') ?? '10');
  const message = url.searchParams.get('message');

  const qs = new URLSearchParams();
  qs.set('page', page.toString());
  qs.set('limit', limit.toString());

  try {
    const res = await fetch(`/api/v1/posts?${qs.toString()}`, {
      headers: {
        cookie: request.headers.get('cookie') as string,
      },
    });

    const jsonResponse = await res.json();

    return { posts: jsonResponse.data, message: message && JSON.parse(message) };
  } catch (_) {
    return fail(HTTP_CODE_BAD_REQUEST, {
      status: API_STATUS.FAIL,
      message: MESSAGE_POST_FETCH_ERROR,
    });
  }
};
