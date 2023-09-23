import { dev } from '$app/environment';
import { BACKEND_URL } from '$env/static/private';
import type { HandleFetch } from '@sveltejs/kit';

export const handleFetch: HandleFetch = async ({ event, request, fetch }) => {
  if (request.url.includes('/api') && !dev) {
    request = new Request(request.url.replace(event.url.origin, BACKEND_URL), request);
  }

  return fetch(request);
};
