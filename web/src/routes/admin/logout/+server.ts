import { HTTP_CODE_SEE_OTHER } from '$lib/constants';
import { redirect, type RequestHandler } from '@sveltejs/kit';
import cookie from 'cookie';

export const GET: RequestHandler = async ({ fetch, cookies }) => {
  const response = await fetch('/api/v1/auth/logout', {
    method: 'POST',
    credentials: 'include',
  });

  const apiCookie = cookie.parse(response.headers.get('set-cookie') as string);

  cookies.set('session_token', apiCookie.session_token, {
    path: apiCookie.path,
    expires: new Date(apiCookie.expires),
    secure: false,
    httpOnly: true,
  });

  // TODO: Handle errors

  redirect(HTTP_CODE_SEE_OTHER, '/');
};
