<script lang="ts">
  // Libraries
  import { Modal } from 'flowbite-svelte';
  import EventView from './EventView.svelte';

  // Types
  import type { Event } from '$lib/types';

  // Utils
  import { formatDateObject, isSameDay } from '$lib/utils/date';
  import { invalidateAll } from '$app/navigation';

  // Constants
  import { API_STATUS, QUICK_LOG_NOTES } from '$lib/constants';

  export let events: Event[] = [];

  const today = new Date();
  const id = formatDateObject(today, 'YYYYMMDD');
  $: createdEvent = events.find((event: Event) =>
    isSameDay(event.start, today.toLocaleDateString())
  );
  const event: Event = {
    id: `quick-log-${id}`,
    title: `Work Event #${id}`,
    start: today.toString(),
    end: today.toString(),
    allDay: true,
    type: 'work',
    notes: QUICK_LOG_NOTES,
  };

  let openModal = false;

  function handleCloseModal() {
    openModal = false;
  }

  async function handleCreateEvent() {
    const response = await fetch(`/events`, {
      method: 'POST',
      body: JSON.stringify(event),
    });

    const responseData = await response.json();

    if (responseData.status === API_STATUS.SUCCESS) {
      await invalidateAll();
      openModal = false;
    }

    // TODO: Handle error
    // TODO: Add success feedback e.g. alert
  }
</script>

{#if Boolean(createdEvent)}
  <div
    class="max-w-sm mt-10 space-y-3 p-6 bg-white border border-zinc-200 rounded-lg shadow dark:bg-zinc-900 dark:border-zinc-800"
  >
    <h5 class="text-2xl font-bold font-sans-pro">{createdEvent?.title}</h5>
    <p class="font-normal">You have already logged your work for today! 🎉🎉🎉</p>
  </div>
{:else}
  <div
    class="max-w-sm mt-10 space-y-3 p-6 bg-white border border-zinc-200 rounded-lg shadow dark:bg-zinc-900 dark:border-zinc-800"
  >
    <h5 class="text-2xl font-bold font-sans-pro">Have you worked today?</h5>
    <p class="font-normal">Create a quick work log for today.</p>
    <div class="action">
      <button class="inline-block btn" on:click={() => (openModal = true)}>Quick Log</button>
    </div>
  </div>
{/if}

<Modal
  class="dark:bg-zinc-800"
  title=""
  size="sm"
  bind:open={openModal}
  on:close={handleCloseModal}
>
  <EventView activeEvent={event} hasActions={false} />
  <div class="actions">
    <button class="btn w-full" on:click={handleCreateEvent}>Create</button>
  </div>
</Modal>