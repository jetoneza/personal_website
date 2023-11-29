// Types
import type { PageServerLoad } from './$types';

// Constants
import { API_STATUS } from '$lib/constants';

export const load: PageServerLoad = async ({ request, fetch }) => {
  try {
    const res = await fetch(`/api/v1/events?filter=month`, {
      headers: {
        cookie: request.headers.get('cookie') as string,
      },
    });

    const jsonResponse = await res.json();

    if (jsonResponse.status === API_STATUS.FAIL) {
      throw jsonResponse;
    }

    return { events: jsonResponse.data };
  } catch (_) {
    return { events: [] };
  }
};
