<script lang="ts">
  import { enhance } from '$app/forms';
  import Input from '$lib/components/Input.svelte';
  import { API_STATUS } from '$lib/constants';
  import type { ActionData } from './$types';

  export let form: ActionData;
</script>

<svelte:head>
  <title>Login</title>
</svelte:head>

<div class="container login mx-auto p-6 flex flex-col justify-center items-center h-screen">
  <div class="title font-bold text-3xl font-sans-pro">Login to Dashboard</div>
  <div
    class="mt-6 form-wrapper w-80 p-6 rounded-lg drop-shadow-lg bg-white dark:bg-transparent dark:border-2 dark:border-zinc-700"
  >
    {#if form?.status === API_STATUS.FAIL}
      <p class="p-4 bg-pink-100 border border-red-500 rounded-lg text-red-500 my-2 text-sm">
        {form?.message}
      </p>
    {/if}

    <form class="flex flex-col space-y-4 mt-4" method="POST" use:enhance>
      <Input type="email" name="email" label="Email" required={true} />
      <Input type="password" name="password" label="Password" required={true} />
      <button type="submit" class="btn"> Submit </button>
    </form>
  </div>
</div>
