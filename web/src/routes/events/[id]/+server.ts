import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

// TODO: Use a cleaner approach for this endpoint
export const DELETE: RequestHandler = async ({ params, fetch }) => {
  const { id } = params;

  const response = await fetch(`/api/v1/events/${id}`, {
    method: 'DELETE',
  });

  const responseJson = await response.json()

  return json(responseJson)
};
