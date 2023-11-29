<script lang="ts">
  // Types
  import type { Event } from '$lib/types';
  import { formatDate } from '$lib/utils/date';
  import type { PageData } from './$types';

  // Props
  export let data: PageData;

  $: daysWorked = data.events.reduce(
    (acc: number, curr: Event) => (curr.type === 'work' ? acc + 1 : acc),
    0
  );
</script>

<main class="blog-dashboard">
  <div class="date absolute top-6 right-6 font-semibold">Today is {formatDate(new Date().toString())}</div>
  <div class="actions w-full flex justify-between items-center">
    <h1 class="font-bold text-2xl font-sans-pro">Welcome, Jet</h1>
  </div>

  <div
    class="max-w-sm mt-10 space-y-5 p-6 bg-white border border-zinc-200 rounded-lg shadow dark:bg-zinc-900 dark:border-zinc-800"
  >
    <div class="font-semibold">Days worked this month</div>
    <div class="text-6xl font-bold font-sans-pro">{daysWorked}</div>
  </div>

  <div
    class="max-w-sm mt-10 space-y-3 p-6 bg-white border border-zinc-200 rounded-lg shadow dark:bg-zinc-900 dark:border-zinc-800"
  >
    <h5 class="text-2xl font-bold font-sans-pro">Post for your blog</h5>
    <p class="font-normal">Add a post to your blog and share it to your friends.</p>
    <div class="action">
      <a href="/admin/dashboard/blog/create" class="inline-block btn"> Create Post </a>
    </div>
  </div>
</main>
