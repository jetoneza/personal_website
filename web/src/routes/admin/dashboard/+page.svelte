<script lang="ts">
  // Components
  import QuickWorkLog from '$lib/components/QuickWorkLog.svelte';
  import RateCalculator from '$lib/components/RateCalculator.svelte';

  // Utils
  import { formatDate } from '$lib/utils/date';

  // Types
  import type { PageData } from './$types';

  // Props
  export let data: PageData;
</script>

<main class="blog-dashboard">
  <div class="date absolute top-6 right-6 font-semibold">
    Today is {formatDate(new Date().toString())}
  </div>
  <div class="actions w-full flex justify-between items-center">
    <h1 class="font-bold text-2xl font-sans-pro">Welcome, Jet</h1>
  </div>

  <RateCalculator events={data.events} />

  <QuickWorkLog events={data.events} />

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
