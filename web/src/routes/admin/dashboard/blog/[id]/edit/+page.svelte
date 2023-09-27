<script lang="ts">
  import { enhance } from '$app/forms';
  import PostForm from '$lib/components/forms/PostForm.svelte';
  import { API_STATUS } from '$lib/constants';
  import type { ActionData, PageData } from './$types';

  export let form: ActionData;
  export let data: PageData;

  const {
    title,
    slug,
    description,
    category,
    meta_title,
    meta_description,
    meta_keyword,
    content,
  } = data.post;
</script>

<div>
  {#if form?.status === API_STATUS.FAIL}
    <p class="p-4 bg-pink-100 border border-red-500 rounded-lg text-red-500 my-6 text-sm">
      {form?.message}
    </p>
  {/if}
  <form class="w-full" method="POST" use:enhance>
    <PostForm
      action="Edit Post"
      {title}
      {slug}
      {description}
      {category}
      {meta_title}
      {meta_description}
      {meta_keyword}
      {content}
    />
  </form>
</div>
