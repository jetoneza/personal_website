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
    content,
    meta_title,
    meta_description,
    meta_keyword,
  } = data.post;
</script>

<div class="container flex flex-col mx-auto gap-8">
  <h1 class="font-bold text-xl">Edit Post</h1>

  <div>
    {#if form?.status === API_STATUS.FAIL}
      <p class="p-4 bg-pink-100 border border-red-500 rounded-lg text-red-500 my-6 text-sm">
        {form?.message}
      </p>
    {/if}
    {#if form?.status === API_STATUS.SUCCESS}
      <p class="p-4 bg-green-100 border border-green-700 rounded-lg text-green-700 my-6 text-sm">
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
</div>
