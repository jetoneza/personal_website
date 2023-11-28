<script lang="ts">
  import type { Event } from '$lib/types';
  import { formatDate } from '$lib/utils/date';

  export let activeEvent: Event;
</script>

<div class="flex flex-col gap-2">
  <div class="title flex flex-col gap-1">
    <span class="text-gray-700 font-bold text-2xl dark:text-white">{activeEvent.title}</span>
    <span class="event-date text-sm text-gray-700 dark:text-white">
      {formatDate(activeEvent.start)}
    </span>
  </div>

  <div class="notes text-sm text-gray-600 dark:text-white pt-4">
    {#if activeEvent.notes}
      {activeEvent.notes}
    {:else}
      <p class="italic">No notes provided.</p>
    {/if}
  </div>
</div>
