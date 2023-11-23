import { fail, redirect } from '@sveltejs/kit';

// Types
import type { PageServerLoad } from './$types';
import type { Actions } from './$types';

// Constants
import { API_STATUS, HTTP_CODE_BAD_REQUEST, MESSAGE_EVENT_CREATION_ERROR } from '$lib/constants';

export const load: PageServerLoad = async () => {
  // TODO: Remove static data and use real data
  const events = [
    {
      id: 2,
      start: new Date('11-21-2023'),
      end: new Date('11-21-2023'),
      allDay: true,
      notes: 'Chill day at work.',
      type: 'work',
      createdAt: new Date(),
    },
    {
      id: 4,
      start: new Date('11-21-2023'),
      end: new Date('11-24-2023'),
      allDay: true,
      notes: 'Project B - Implementation',
      type: 'task',
      createdAt: new Date(),
    },
    {
      id: 3,
      start: new Date('12-25-2023'),
      end: new Date('12-25-2023'),
      allDay: true,
      notes: 'Christmas Day 🌲🎁🎉',
      type: 'holiday',
      createdAt: new Date(),
    },
  ];

  try {
    return { events };
  } catch (_) {
    return { events: [] };
  }
};

export const actions: Actions = {
  default: async ({ request }) => {
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
