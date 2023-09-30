<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;
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

<div class="flex flex-col gap-16 md:gap-8">
  <div class="flex flex-col gap-2">
    <h1 class="font-bold text-4xl font-sans-pro">Projects</h1>
    <p class="text-secondary animate-in">
      Check out the projects I enjoy working on. Link to my Github account is <a
        href="https://github.com/jetoneza"
        class="font-bold">here.</a
      >
    </p>
  </div>
  <ul class="divide-y divide-gray-200 dark:divide-gray-700">
    <li class="py-8">
      <article>
        <div class="space-y-2">
          <div class="space-y-3">
            <div class="space-y-4">
              <div>
                <h2 class="text-xl font-bold font-sans-pro leading-8 tracking-tight">
                  <a href="/blog/building-a-cycling-trainer-using-rust">Cycling Trainer</a>
                </h2>
              </div>
              <div class="prose max-w-none dark:prose-invert">
                Cycling Trainer App project, crafted with Rust and integrated with Bluetooth Low
                Energy. This uses Tauri, a library that lets you create desktop apps using Rust for
                the backend and Svelte for the frontend.
                <a
                  href="/blog/building-a-cycling-trainer-using-rust"
                  class="text-cyan-500 hover:text-cyan-600 dark:hover:text-cyan-400 no-underline"
                >
                  Read more.
                </a>
              </div>
              <img
                alt="Cycling Trainer"
                src="https://res.cloudinary.com/dj0l6hcwu/image/upload/c_scale,w_700/v1696052264/personal_website/cycling_trainer_workout_qo1rlq.png"
              />
            </div>
          </div>
        </div>
      </article>
    </li>
    <li class="py-8">
      <article>
        <div class="space-y-2">
          <div class="space-y-3">
            <div class="space-y-4">
              <div>
                <h2 class="text-xl font-bold font-sans-pro leading-8 tracking-tight">
                  <a href="/blog/building-my-personal-website">Personal Website</a>
                </h2>
              </div>
              <div class="prose max-w-none dark:prose-invert">
                This website is built using Go and Svelte Kit.
                <a
                  href="/blog/building-my-personal-website"
                  class="text-cyan-500 hover:text-cyan-600 dark:hover:text-cyan-400 no-underline"
                >
                  Read more.
                </a>
              </div>
              <img
                alt="Site Design"
                src="https://res.cloudinary.com/dj0l6hcwu/image/upload/c_scale,w_700/v1696086546/personal_website/go_sveltekit_design_hivwz3.png"
                class="rounded-lg mt-10 bg-white drop-shadow-2xl"
              />
            </div>
          </div>
        </div>
      </article>
    </li>
  </ul>
</div>
