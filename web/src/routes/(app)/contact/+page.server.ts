import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
  // TODO: Add image
  const meta = {
    title: 'Contact | Jet Ordaneza',
    description: `If you'd like to get in touch, you can email me at jet.oneza@gmail.com or check out my linkedin and github profiles`,
    keywords: 'software engineer,reactjs,react,nodejs,node,rust,go,golang,tailwindcss',
    url: 'https://www.jetordaneza.com/contact',
    imageUrl: '',
  }

  return { meta };
};
