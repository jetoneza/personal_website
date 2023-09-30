import { dev } from '$app/environment';
import {
  type Handle,
  type HandleFetch,
  type Redirect,
  type HttpError,
  type RequestEvent,
  error,
  redirect,
} from '@sveltejs/kit';
import { BACKEND_URL } from '$env/static/private';
import { HTTP_CODE_NOT_FOUND, HTTP_CODE_SEE_OTHER } from '$lib/constants';

const protectedRoutes = ['/admin/dashboard'];

export const handle: Handle = async ({ event, resolve }) => {
  const sessionAction = checkSession(event);

  if (sessionAction) {
    throw sessionAction;
  }

  const theme = checkTheme(event);

  if (theme) {
    return await resolve(event, {
      transformPageChunk: ({ html }) => html.replace('class=""', `class="${theme}"`),
    });
  }

  return resolve(event);
};

export const handleFetch: HandleFetch = async ({ event, request, fetch }) => {
  if (request.url.includes('/api') && !dev) {
    request = new Request(request.url.replace(event.url.origin, BACKEND_URL), request);
  }

  return fetch(request);
};

const checkSession = (event: RequestEvent) => {
  const sessionToken = event.cookies.get('session_token');
  const hasSessionEnded = sessionToken === 'session-ended' || sessionToken === 'expired';

  let action: HttpError | Redirect | null = null;

  const isProtected = checkProtectedRoute(event.url.pathname);

  if ((hasSessionEnded || !sessionToken) && isProtected) {
    action = error(HTTP_CODE_NOT_FOUND, {
      message: 'Not found',
    });
  }

  if (sessionToken && !hasSessionEnded && event.url.pathname.includes('/login')) {
    action = redirect(HTTP_CODE_SEE_OTHER, '/admin/dashboard');
  }

  return action;
};

const checkTheme = (event: RequestEvent): string | null => {
  let theme = 'dark';

  const newTheme = event.url.searchParams.get('theme');
  const cookieTheme = event.cookies.get('theme');

  if (newTheme) {
    theme = newTheme;
  } else if (cookieTheme) {
    theme = cookieTheme;
  }

  return theme;
};

const checkProtectedRoute = (path: string): boolean => {
  for (const route of protectedRoutes) {
    if (path.includes(route)) {
      return true;
    }
  }

  return false;
};
