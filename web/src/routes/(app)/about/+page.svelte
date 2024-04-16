<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;

  const experiences = [
    {
      time: '2023 - Present',
      title: 'Software Engineer',
      company: 'Bookipi',
      link: 'https://bookipi.com',
    },
    {
      time: '2018 - 2023',
      title: 'Senior Software Engineer',
      company: 'Poplar Homes',
      link: 'https://poplarhomes.com',
    },
    {
      time: '2016 - 2018',
      title: 'Software Engineer',
      company: 'Satoshi Citadel Industries',
    },
    {
      time: '2015 - 2016',
      title: 'Web Developer',
      company: 'Twidl Inc.',
    },
    {
      time: '2014 - 2015',
      title: 'Web Developer',
      company: 'Freelance',
    },
  ];
</script>

<svelte:head>
  <title>{data.meta.title}</title>

  <meta name="description" content={data.meta.description} />
  <meta name="keywords" content={data.meta.keywords} />

  <!-- Social media -->
  <meta name="twitter:card" content="summary" />
  <meta name="twitter:title" content={data.meta.title} />
  <meta name="twitter:description" content={data.meta.description} />
  <meta name="twitter:image" content={data.meta.imageUrl} />
  <meta name="twitter:url" content={data.meta.url} />

  <meta property="og:title" content={data.meta.title} />
  <meta property="og:description" content={data.meta.description} />
  <meta property="og:image" content={data.meta.imageUrl} />
  <meta property="og:url" content={data.meta.url} />
  <meta property="og:type" content="website" />
</svelte:head>

<h1 class="font-bold text-4xl font-sans-pro">About Me</h1>
<div class="dark:text-slate-200">Hello! I'm Jet Ordaneza.</div>
<div class="space-y-10 mt-10">
  <div class="content space-y-4 dark:text-slate-200">
    <p class="text-justify">
      When I'm not coding, you can find me doing a few different things. I enjoy playing games,
      riding my bicycle, and strumming my guitar. These are my go-to hobbies that help me relax and
      have fun.
    </p>
  </div>

  <div class="experience space-y-4 mt-8 dark:text-slate-200">
    <h1 class="font-bold text-2xl dark:text-white font-sans-pro">Experience</h1>

    <p class="text-justify">
      I’ve had the privilege of working in diverse tech environments, where I’ve honed my skills and
      tackled a variety of exciting challenges.
    </p>

    <p class="text-justify">
      My expertise spans a range of technologies, including TypeScript, Node, React, GraphQL,
      Tailwind CSS, Svelte, Go, and Rust.
    </p>
  </div>

  <ul class="flex flex-col gap-4">
    {#each experiences as experience}
      <li class="flex flex-col-reverse min-[425px]:flex-row min-[425px]:gap-10">
        <div class="time text-sm w-[115px] min-[425px]:text-base font-bold">{experience.time}</div>
          <a href="{experience.link}" target="_blank">
            <div class="title font-bold text-cyan-500">{experience.title}</div>
          <div class="company text-sm">{experience.company}</div>
        </a>
      </li>
    {/each}
  </ul>

  <div class="links space-y-4 mt-8 dark:text-slate-200">
    <h1 class="font-bold text-2xl dark:text-white font-sans-pro">Links</h1>

    <div class="flex gap-4">
      <a
        href="https://linkedin.com/in/jetoneza"
        target="_blank"
        class="flex items-center p-3 rounded-lg border border-slate-800 dark:border-zinc-700 gap-2 justify-center w-36 font-bold hover:bg-[#0077B5] hover:text-white hover:border-[#0077B5]"
      >
        <span>
          <svg xmlns="http://www.w3.org/2000/svg" height="1em" viewBox="0 0 448 512">
            <path
              fill="currentColor"
              d="M100.28 448H7.4V148.9h92.88zM53.79 108.1C24.09 108.1 0 83.5 0
                53.8a53.79 53.79 0 0 1 107.58 0c0 29.7-24.1 54.3-53.79 54.3zM447.9
                448h-92.68V302.4c0-34.7-.7-79.2-48.29-79.2-48.29 0-55.69 37.7-55.69
                76.7V448h-92.78V148.9h89.08v40.8h1.3c12.4-23.5 42.69-48.3 87.88-48.3
                94 0 111.28 61.9 111.28 142.3V448z"
            />
          </svg>
        </span>
        <span class="hidden min-[425px]:block">LinkedIn</span>
      </a>

      <a
        href="https://github.com/jetoneza"
        target="_blank"
        class="flex items-center p-3 rounded-lg border border-slate-800 dark:border-zinc-700
        gap-2 justify-center w-36 font-bold hover:bg-[#333] hover:border-[#333 hover:text-[#f5f5f5]
        dark:hover:text-[#333] dark:hover:bg-[#f5f5f5] dark:hover:border-[#f5f5f5]"
      >
        <span>
          <svg height="1.5em" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
            <title>GitHub</title>
            <path
              fill="currentColor"
              d="M12.001 2C6.47598 2 2.00098 6.475 2.00098 12C2.00098 16.425 4.86348
                20.1625 8.83848 21.4875C9.33848 21.575 9.52598 21.275 9.52598 21.0125C9.52598
                20.775 9.51348 19.9875 9.51348 19.15C7.00098 19.6125 6.35098 18.5375 6.15098
                17.975C6.03848 17.6875 5.55098 16.8 5.12598 16.5625C4.77598 16.375 4.27598
                15.9125 5.11348 15.9C5.90098 15.8875 6.46348 16.625 6.65098 16.925C7.55098
                18.4375 8.98848 18.0125 9.56348 17.75C9.65098 17.1 9.91348 16.6625 10.201
                16.4125C7.97598 16.1625 5.65098 15.3 5.65098 11.475C5.65098 10.3875 6.03848
                9.4875 6.67598 8.7875C6.57598 8.5375 6.22598 7.5125 6.77598 6.1375C6.77598
                6.1375 7.61348 5.875 9.52598 7.1625C10.326 6.9375 11.176 6.825 12.026
                6.825C12.876 6.825 13.726 6.9375 14.526 7.1625C16.4385 5.8625 17.276
                6.1375 17.276 6.1375C17.826 7.5125 17.476 8.5375 17.376 8.7875C18.0135
                9.4875 18.401 10.375 18.401 11.475C18.401 15.3125 16.0635 16.1625
                13.8385 16.4125C14.201 16.725 14.5135 17.325 14.5135 18.2625C14.5135
                19.6 14.501 20.675 14.501 21.0125C14.501 21.275 14.6885 21.5875
                15.1885 21.4875C19.259 20.1133 21.9999 16.2963 22.001 12C22.001
                6.475 17.526 2 12.001 2Z"
            />
          </svg>
        </span>
        <span class="hidden min-[425px]:block">Github</span>
      </a>

      <a
        href="https://www.youtube.com/@jetrooper_gaming"
        target="_blank"
        class="flex gap-2 items-center p-3 rounded-lg border border-slate-800 dark:border-zinc-700 justify-center w-36 font-bold hover:bg-[#FF0000] hover:border-[#FF0000] hover:text-white"
      >
        <span>
          <svg xmlns="http://www.w3.org/2000/svg" height="1.2em" viewBox="0 0 576 512">
            <path
              fill="currentColor"
              d="M549.655 124.083c-6.281-23.65-24.787-42.276-48.284-48.597C458.781
            64 288 64 288 64S117.22 64 74.629 75.486c-23.497 6.322-42.003 24.947-48.284
            48.597-11.412 42.867-11.412 132.305-11.412 132.305s0 89.438 11.412
            132.305c6.281 23.65 24.787 41.5 48.284 47.821C117.22 448 288 448 288
            448s170.78 0 213.371-11.486c23.497-6.321 42.003-24.171 48.284-47.821
            11.412-42.867 11.412-132.305 11.412-132.305s0-89.438-11.412-132.305zm-317.51
            213.508V175.185l142.739 81.205-142.739 81.201z"
            />
          </svg>
        </span>
        <span class="hidden min-[425px]:block">YouTube</span>
      </a>
    </div>
  </div>
</div>
