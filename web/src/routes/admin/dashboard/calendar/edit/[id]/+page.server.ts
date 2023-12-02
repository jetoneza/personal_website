import {
  API_STATUS,
  HTTP_CODE_BAD_REQUEST,
  HTTP_CODE_SEE_OTHER,
  MESSAGE_EVENT_CREATION_ERROR,
  MESSAGE_EVENT_UPDATE_SUCCESS,
} from '$lib/constants';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, params }) => {
  const { id } = params;

  // TODO: Add error handling
  const res = await fetch(`/api/v1/events/${id}`);
  const jsonResponse = await res.json();
  const event = jsonResponse.data;

  return {
    event,
  };
};

export const actions: Actions = {
  default: async ({ request, fetch, params }) => {
    const data = await request.formData();

    const { id } = params;

    const start = new Date(data.get('start') as string);
    const end = new Date(data.get('end') as string);

    try {
      const response = await fetch(`/api/v1/events/${id}`, {
        method: 'PUT',
        headers: {
          cookie: request.headers.get('cookie') as string,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          title: data.get('title'),
          type: data.get('type'),
          all_day: data.get('all_day'),
          notes: data.get('notes'),
          start,
          end,
        }),
      });

      const result = JSON.parse(await response.text());

      if (result.status === API_STATUS.FAIL) {
        throw new Error(MESSAGE_EVENT_CREATION_ERROR);
      }
    } catch (_) {
      return fail(HTTP_CODE_BAD_REQUEST, {
        status: API_STATUS.FAIL,
        message: MESSAGE_EVENT_CREATION_ERROR,
      });
    }

    throw redirect(
      HTTP_CODE_SEE_OTHER,
      `/admin/dashboard/calendar?message={"type":"success","value":"${MESSAGE_EVENT_UPDATE_SUCCESS}"}`
    );
  },
};
