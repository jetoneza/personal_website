<script lang="ts">
  import { enhance } from '$app/forms';
  import EventForm from '$lib/components/forms/EventForm.svelte';
</script>

<div class="container flex flex-col mx-auto gap-8">
  <h1 class="font-bold text-xl">Edit Event</h1>
  <form class="w-full" method="POST" use:enhance>
    <EventForm action="Edit" />
  </form>
</div>
