import { HTTP_CODE_SEE_OTHER } from '$lib/constants';
import { redirect, type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ fetch }) => {
  await fetch('/api/v1/auth/logout', {
    method: 'POST',
  });

  // TODO: Handle errors

  throw redirect(HTTP_CODE_SEE_OTHER, '/');
};
