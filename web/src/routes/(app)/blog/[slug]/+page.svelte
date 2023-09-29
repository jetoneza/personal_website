<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<svelte:head>
  <title>{data.post.meta_title}</title>
  <meta name="description" content={data.post.meta_description} />
  <meta name="keywords" content={data.post.meta_keywords} />

  <!-- Social media -->
  <meta name="twitter:card" content="summary"/>
  <meta name="twitter:title" content={data.post.meta_title} />
  <meta name="twitter:description" content={data.post.meta_description} />
  <meta name="twitter:image" content={data.post.meta_image_url } />
  <meta name="twitter:url" content="https://www.jetordaneza.com/blog/{data.post.slug}" />

  <meta property="og:title" content={data.post.meta_title} />
  <meta property="og:description" content={data.post.description} />
  <meta property="og:image" content={data.post.meta_image_url } />
  <meta property="og:url" content="https://www.jetordaneza.com/blog/{data.post.slug}" />
  <meta property="og:type" content="blog" />
</svelte:head>

<div class="title-wrapper flex flex-col space-y-4">
  <h1 class="font-extrabold text-4xl">
    {data.post.title}
  </h1>
  <time class="text-sm font-bold text-cyan-500" datetime={data.post.createdAt}>
    {data.post.formattedCreatedAt}
  </time>
</div>

<article class="prose dark:prose-invert mt-10">
  {@html data.post.content}
</article>
