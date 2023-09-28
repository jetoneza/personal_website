<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<div class="title-wrapper flex flex-col space-y-4">
  <h1 class="font-extrabold text-4xl"> 
    {data.post.title}
  </h1> 
  <time class="text-sm font-bold text-cyan-500" datetime={data.post.createdAt}>{data.post.formattedCreatedAt}</time>
</div>

<article class="prose dark:prose-invert mt-10">
  {@html data.post.content}
</article>
