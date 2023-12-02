<script lang="ts">
  // Common Components
  import Input from '$lib/components/Input.svelte';

  export let action = 'Create';
  export let title = '';
  export let type = '';
  export let start = '';
  export let end = '';
  export let all_day = false;
  export let notes = '';
</script>

<div class="flex flex-col gap-4">
  <!-- TODO: Support time selection -->
  <Input type="text" name="title" label="Title" placeholder="(No Title)" value={title} required />
  <Input
    type="text"
    name="type"
    label="Type"
    value={type}
    placeholder="Enter type e.g. work, task"
    required
  />
  <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
    <!-- TODO: Add date checking. Start should be less than end -->
    <Input type="date" name="start" value={start} label="Start" />
    <Input type="date" name="end" value={end} label="End" />
  </div>
  <div class="input-wrapper">
    <label for="all_day" class="block mb-2 text-md font-bold dark:text-white"> All Day? </label>
    <select
      value={all_day}
      id="all_day"
      name="all_day"
      class="
        h-10 w-full rounded-md border border-input bg-background px-3
        py-2 text-sm outline-none focus:outline-zinc-500 dark:bg-zinc-800 dark:border-zinc-700
        dark:focus:outline-zinc-600
      "
    >
      <option selected value={false}>No</option>
      <option value={true}>Yes</option>
    </select>
  </div>
  <div class="textarea-wrapper mt-2">
    <label for="content" class="block mb-2 text-md font-bold dark:text-white">Content</label>
    <textarea
      value={notes}
      id="notes"
      name="notes"
      rows="10"
      class="
        w-full rounded-md border border-input bg-background px-3 py-2
        text-sm outline-none focus:outline-zinc-500 font-mono text-black dark:bg-zinc-800
        dark:border-zinc-700 dark:text-white
      "
    />
  </div>
  <button class="btn mt-2">{action}</button>
</div>
