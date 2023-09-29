<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<svelte:head>
  <title>{data.meta.title}</title>

  <meta name="description" content={data.meta.description} />
  <meta name="keywords" content={data.meta.keywords} />

  <!-- Social media -->
  <meta name="twitter:card" content="summary"/>
  <meta name="twitter:title" content={data.meta.title} />
  <meta name="twitter:description" content={data.meta.description} />
  <meta name="twitter:image" content={data.meta.imageUrl } />
  <meta name="twitter:url" content={data.meta.url} />

  <meta property="og:title" content={data.meta.title} />
  <meta property="og:description" content={data.meta.description} />
  <meta property="og:image" content={data.meta.imageUrl } />
  <meta property="og:url" content={data.meta.url} />
  <meta property="og:type" content="website" />
</svelte:head>

<div class="contact">
  <h1 class="font-bold text-2xl">Contact Me</h1>
  <p class="mt-10">
    If you'd like to get in touch, you can email me at <a
      class="font-bold text-cyan-500 underline"
      href="mailto:jet.oneza@gmail.com"
    >
      jet.oneza@gmail.com
    </a>
  </p>
</div>
