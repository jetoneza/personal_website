import type { PageServerLoad } from './$types';

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
