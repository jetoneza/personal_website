import { dev } from '$app/environment';
import { type Handle, type HandleFetch, error } from '@sveltejs/kit';
import { BACKEND_URL } from '$env/static/private';
import { HTTP_CODE_NOT_FOUND } from '$lib/constants';

const protectedRoutes = ['/admin/dashboard'];

export const handle: Handle = async ({ event, resolve }) => {
  const sessionToken = event.cookies.get('session_token');

  if (!sessionToken && protectedRoutes.includes(event.url.pathname)) {
    throw error(HTTP_CODE_NOT_FOUND, {
      message: 'Not found',
    });
  }

  // TODO: fetch user info and load user to locals

  return resolve(event);
};

export const handleFetch: HandleFetch = async ({ event, request, fetch }) => {
  if (request.url.includes('/api') && !dev) {
    request = new Request(request.url.replace(event.url.origin, BACKEND_URL), request);
  }

  return fetch(request);
};
