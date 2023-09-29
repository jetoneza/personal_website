import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
  // TODO: Add image
  const meta = {
    title: 'About | Jet Ordaneza',
    description: `When I'm not coding, you can find me doing a few different things.
      I enjoy playing games, riding my bicycle, and strumming my guitar.
      These are my go-to hobbies that help me relax and have fun.`,
    keywords: 'software engineer,reactjs,react,nodejs,node,rust,go,golang,tailwindcss',
    url: 'https://www.jetordaneza.com/about',
    imageUrl: '',
  };

  return { meta };
};
