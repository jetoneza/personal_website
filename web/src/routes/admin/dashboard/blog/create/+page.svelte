<script lang="ts">
  import { enhance } from '$app/forms';
  import PostForm from '$lib/components/forms/PostForm.svelte';
  import { API_STATUS } from '$lib/constants';
  import type { ActionData } from './$types';

  export let form: ActionData;
</script>

<div class="container flex flex-col mx-auto gap-8">
  <h1 class="font-bold text-xl">Add new Post</h1>

  <div>
    {#if form?.status === API_STATUS.FAIL}
      <p class="p-4 bg-pink-100 border border-red-500 rounded-lg text-red-500 my-6 text-sm">
        {form?.message}
      </p>
    {/if}
    <form class="w-full" method="POST" use:enhance>
      <PostForm action="Create Post" />
    </form>
  </div>
</div>
