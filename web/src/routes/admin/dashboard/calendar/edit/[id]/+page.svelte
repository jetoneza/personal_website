<script lang="ts">
  // Components
  import EventForm from '$lib/components/forms/EventForm.svelte';

  // Utils
  import { enhance } from '$app/forms';
  import { formatInputDate } from '$lib/utils/date';

  // Types
  import type { PageData } from './$types';

  export let data: PageData;

  const { title, type, all_day, notes, start, end } = data.event;
</script>

<div class="container flex flex-col mx-auto gap-8">
  <h1 class="font-bold text-xl">Edit Event</h1>
  <form class="w-full" method="POST" use:enhance>
    <EventForm
      action="Edit"
      {title}
      {type}
      {notes}
      allDay={all_day}
      start={formatInputDate(new Date(start))}
      end={formatInputDate(new Date(end))}
    />
  </form>
</div>
