<script lang="ts">
  // Components
  import EventForm from '$lib/components/forms/EventForm.svelte';

  // Utils
  import { enhance } from '$app/forms';
  import { formatInputDate } from '$lib/utils/date';

  // Types
  import type { ActionData, PageData } from './$types';

  // Constants
  import { API_STATUS } from '$lib/constants';

  export let form: ActionData;
  export let data: PageData;

  const { title, type, all_day, notes, start, end } = data.event;
</script>

<div class="container flex flex-col mx-auto gap-8">
  <h1 class="font-bold text-xl">Edit Event</h1>
  <div>
    {#if form?.status === API_STATUS.FAIL}
      <p class="p-4 bg-pink-100 border border-red-500 rounded-lg text-red-500 my-6 text-sm">
        {form?.message}
      </p>
    {/if}
    {#if form?.status === API_STATUS.SUCCESS}
      <p class="p-4 bg-green-100 border border-green-700 rounded-lg text-green-700 my-6 text-sm">
        {form?.message}
      </p>
    {/if}
    <form class="w-full" method="POST" use:enhance>
      <EventForm
        action="Edit"
        {title}
        {type}
        {notes}
        {all_day}
        start={formatInputDate(new Date(start))}
        end={formatInputDate(new Date(end))}
      />
    </form>
  </div>
</div>
