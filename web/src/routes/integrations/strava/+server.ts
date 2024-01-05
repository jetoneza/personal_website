import { API_STATUS, HTTP_CODE_BAD_REQUEST, HTTP_CODE_SEE_OTHER } from '$lib/constants';
import { json, type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ fetch, url }) => {
  const code = url.searchParams.get('code');

  const response = await fetch(`/api/v1/integrations/strava?code=${code}`, {
    method: 'GET',
  });

  const result = JSON.parse(await response.text());

  if (result.status === API_STATUS.FAIL) {
    // TODO: Improve eroor handling
    return json({
      code: HTTP_CODE_BAD_REQUEST,
      message: 'Strava integration failed.',
      ...result,
    })
  }

  return json(
    result
  );
};
