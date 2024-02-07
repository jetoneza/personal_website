import { json } from '@sveltejs/kit';

import type { RequestHandler } from './$types';

import { API_STATUS, HTTP_CODE_BAD_REQUEST, MESSAGE_EVENT_CREATION_ERROR } from '$lib/constants';

// TODO: Use a cleaner approach for this endpoint
export const POST: RequestHandler = async ({ request, fetch }) => {
  const { start, end, title, allDay, type, notes } = await request.json();

  try {
    const response = await fetch('/api/v1/events', {
      method: 'POST',
      headers: {
        cookie: request.headers.get('cookie') as string,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        start: new Date(start),
        end: new Date(end),
        title,
        all_day: allDay.toString(),
        type,
        notes,
      }),
    });

    const result = JSON.parse(await response.text());

    if (result.status === API_STATUS.FAIL) {
      throw new Error('Event creation failed');
    }

    return json({
      status: API_STATUS.SUCCESS,
    });
  } catch (_) {
    return json({
      code: HTTP_CODE_BAD_REQUEST,
      status: API_STATUS.FAIL,
      message: MESSAGE_EVENT_CREATION_ERROR,
    });
  }
};
