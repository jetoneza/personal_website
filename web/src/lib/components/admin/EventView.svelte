<script lang="ts">
  // Libraries
  import { marked } from 'marked';

  import TrashBinSolid from 'flowbite-svelte-icons/TrashBinSolid.svelte';
  import PenSolid from 'flowbite-svelte-icons/PenSolid.svelte';

  // Utils
  import { formatDate } from '$lib/utils/date';

  // Types
  import type { Event } from '$lib/types';

  export let activeEvent: Event;
  export let onDeleteClick: (id: string | number) => void;
</script>

<div class="flex flex-col gap-2 text-slate-800 dark:text-white pt-6">
  <div class="actions absolute top-3 right-14 flex">
    <a
      href="/admin/dashboard/calendar/edit/{activeEvent.id}"
      class="inline-block p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-zinc-700"
    >
      <PenSolid size="md" />
    </a>
    <button
      class="p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-zinc-700"
      on:click={() => {
        onDeleteClick(activeEvent.id);
      }}
    >
      <TrashBinSolid size="md" />
    </button>
  </div>
  <div class="title flex flex-col gap-1">
    <div class="header">
      <span class="font-bold text-2xl">{activeEvent.title}</span>
    </div>
    <span class="event-date text-sm">
      {formatDate(activeEvent.start)}
    </span>
  </div>

  <div class="notes text-sm pt-4 prose dark:prose-invert">
    {#if activeEvent.notes}
      {@html marked(activeEvent.notes)}
    {:else}
      <p class="italic">No notes provided.</p>
    {/if}
  </div>
</div>
