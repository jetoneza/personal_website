<script lang="ts">
  import { enhance } from '$app/forms'
  import Input from '$lib/components/Input.svelte';
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

<div class="contact">
  <h1 class="font-bold text-2xl">Contact Me</h1>
  <p class="mt-10">
    If you'd like to get in touch, you can contact me here!
  </p>

  <form class="contact-form mt-10 space-y-4" method="POST" use:enhance>
    <Input name="name" label="Name" type="text" placeholder="" />
    <Input name="email" label="Email" type="email" placeholder="" />
    <div class="textarea-wrapper mt-2">
      <label for="message" class="block mb-2 text-md font-bold dark:text-white">Message</label>
      <textarea
        value=""
        name="message"
        rows="10"
        class="
          w-full rounded-md border border-input bg-background px-3 py-2
          text-sm outline-none focus:outline-zinc-500 font-mono text-black dark:bg-zinc-800
          dark:border-zinc-700 dark:text-white
        "
    />
    </div>
    <button type="submit" class="btn">Submit</button>
  </form>
</div>
