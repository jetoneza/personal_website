import { fail } from '@sveltejs/kit';

// Types
import type { PageServerLoad } from './$types';
import type { Actions } from './$types';

// Constants
import { API_STATUS, HTTP_CODE_BAD_REQUEST, MESSAGE_EVENT_CREATION_ERROR } from '$lib/constants';

export const load: PageServerLoad = async ({ request, fetch, url }) => {
  const page = Number(url.searchParams.get('page') ?? '1');
  const limit = Number(url.searchParams.get('limit') ?? '10');
  const message = url.searchParams.get('message');

  const qs = new URLSearchParams();
  qs.set('page', page.toString());
  qs.set('limit', limit.toString());

  try {
    const res = await fetch(`/api/v1/events?${qs.toString()}`, {
      headers: {
        cookie: request.headers.get('cookie') as string,
      },
    });

    const jsonResponse = await res.json();

    if (jsonResponse.status === API_STATUS.FAIL) {
      throw jsonResponse;
    }

    return { events: jsonResponse.data, message: message && JSON.parse(message) };
  } catch (_) {
    return { events: [], message: { type: API_STATUS.FAIL, value: MESSAGE_EVENT_CREATION_ERROR } };
  }
};

export const actions: Actions = {
  default: async ({ request, fetch }) => {
    const data = await request.formData();

    const start = new Date(data.get('start') as string);
    const end = new Date(data.get('end') as string);

    try {
      const response = await fetch('/api/v1/events', {
        method: 'POST',
        headers: {
          cookie: request.headers.get('cookie') as string,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          title: data.get('title'),
          start,
          end,
          all_day: data.get('all_day'),
          type: data.get('type'),
        }),
      });

      const result = JSON.parse(await response.text());

      if (result.status === API_STATUS.FAIL) {
        return fail(HTTP_CODE_BAD_REQUEST, {
          ...result,
          message: MESSAGE_EVENT_CREATION_ERROR,
        });
      }

      return {
        status: API_STATUS.SUCCESS,
      };
    } catch (_) {
      return fail(HTTP_CODE_BAD_REQUEST, {
        status: API_STATUS.FAIL,
        message: MESSAGE_EVENT_CREATION_ERROR,
      });
    }
  },
};
