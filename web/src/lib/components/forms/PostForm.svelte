<script lang="ts">
  import Input from '$lib/components/Input.svelte';

  export let action = 'Create Post';
  export let title = '';
  export let slug = '';
  export let description = '';
  export let category = '';
  export let meta_title = '';
  export let meta_description = '';
  export let meta_keywords = '';
  export let meta_image_url = '';
  export let content = '';
  export let published = false;
</script>

<div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
  <div class="flex flex-col gap-2">
    <Input type="text" name="title" value={title} label="Title" required={true} />
    <Input type="text" name="slug" value={slug} label="Slug" />
    <Input type="text" name="description" value={description} label="Description" />
    <Input type="text" name="category" value={category} label="Category" />
    <div class="input-wrapper">
      <label for="published" class="block mb-2 text-md font-bold dark:text-white">
        Publish Post
      </label>
      <select
        value={published}
        id="published"
        name="published"
        class="h-10 w-full rounded-md border border-input bg-background px-3
        py-2 text-sm outline-none focus:outline-zinc-500 dark:bg-zinc-800 dark:border-zinc-700
        dark:focus:outline-zinc-600"
      >
        <option selected value={false}>No</option>
        <option value={true}>Yes</option>
      </select>
    </div>
  </div>
  <div class="flex flex-col gap-2">
    <Input type="text" name="meta_title" value={meta_title} label="Meta Title" />
    <Input type="text" name="meta_description" value={meta_description} label="Meta Description" />
    <Input type="text" name="meta_keywords" value={meta_keywords} label="Meta Keywords" />
    <Input type="text" name="meta_image_url" value={meta_image_url} label="Meta Image URL" />
  </div>
</div>
<div class="textarea-wrapper mt-2">
  <label for="content" class="block mb-2 text-md font-bold dark:text-white">Content</label>
  <textarea
    id="content"
    name="content"
    rows="30"
    value={content}
    class="w-full rounded-md border border-input bg-background px-3 py-2
    text-sm outline-none focus:outline-zinc-500 font-mono text-black dark:bg-zinc-800
    dark:border-zinc-700 dark:text-white"
  />
</div>
<button class="btn mt-2">{action}</button>
