<script lang="ts">
  import { enhance } from '$app/forms';
  import Input from '$lib/components/Input.svelte';
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
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
        <div class="flex flex-col gap-2">
          <Input type="text" name="title" label="Title" required={true} />
          <Input type="text" name="slug" label="Slug" />
          <Input type="text" name="description" label="Description" />
          <Input type="text" name="category" label="Category" />
        </div>
        <div class="flex flex-col gap-2">
          <Input type="text" name="meta_title" label="Meta Title" />
          <Input type="text" name="meta_description" label="Meta Description" />
          <Input type="text" name="meta_keyword" label="Meta Keyword" />
        </div>
      </div>
      <div class="textarea-wrapper mt-2">
        <label for="content" class="block mb-2 text-md font-bold dark:text-white">Content</label>
        <textarea
          id="content"
          name="content"
          rows="30"
          class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm focus:outline-zinc-500 font-mono text-black"
        />
      </div>
      <button class="btn mt-2">Create Post</button>
    </form>
  </div>
</div>
