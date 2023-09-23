import { fail, redirect } from '@sveltejs/kit';

// Types
import type { Actions } from './$types';

// Constants
import {
  API_STATUS,
  HTTP_CODE_BAD_REQUEST,
  HTTP_CODE_SEE_OTHER,
  MESSAGE_INCORRECT_CREDENTIALS,
} from '$lib/constants';

export const actions = {
  default: async ({ request, fetch }) => {
    const data = await request.formData();

    const response = await fetch('/api/v1/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email: data.get('email'),
        password: data.get('password'),
      }),
    });

    const result = JSON.parse(await response.text());

    if (result.status === API_STATUS.FAIL) {
      return fail(HTTP_CODE_BAD_REQUEST, {
        ...result,
        message: MESSAGE_INCORRECT_CREDENTIALS,
      });
    }

    throw redirect(HTTP_CODE_SEE_OTHER, '/admin/dashboard');
  },
} satisfies Actions;
