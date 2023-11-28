<script lang="ts">
  // Utils
  import { formatDate } from '$lib/utils/date';

  // Types
  import type { Event } from '$lib/types';

  export let activeEvent: Event;
</script>

<div class="flex flex-col gap-2 text-slate-800 dark:text-white">
  <div class="title flex flex-col gap-1">
    <span class="font-bold text-2xl">{activeEvent.title}</span>
    <span class="event-date text-sm">
      {formatDate(activeEvent.start)}
    </span>
  </div>

  <div class="notes text-sm pt-4">
    {#if activeEvent.notes}
      {activeEvent.notes}
    {:else}
      <p class="italic">No notes provided.</p>
    {/if}
  </div>
</div>
