<script lang="ts">
  import { enhance } from '$app/forms';
  import type { SubmitFunction } from '@sveltejs/kit';
  import { onMount } from 'svelte';

  let currentTheme = '';

  onMount(() => {
    currentTheme = document.documentElement.getAttribute('class') ?? '';
  });

  const handleUpdateTheme: SubmitFunction = ({ action }) => {
    const theme = action.searchParams.get('theme');

    if (theme) {
      document.documentElement.setAttribute('class', theme);
      currentTheme = theme;
    }
  };
</script>

<form method="POST" use:enhance={handleUpdateTheme}>
  <button
    class="flex {currentTheme === 'dark'
      ? ''
      : 'hidden'} items-center text-slate-800 dark:text-white"
    formaction="/?/setTheme&theme=light"
  >
    <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="w-5 h-5"
      ><path
        d="M12 3.5C7.30558 3.5 3.5 7.30558 3.5 12C3.5 16.6944 7.30558 20.5 12 20.5C16.4176 20.5 20.0476 17.1303 20.4608 12.8207C20.4801 12.6202 20.377 12.4277 20.1995 12.3324C20.0219 12.2372 19.8045 12.2578 19.6481 12.3848C18.7884 13.0824 17.6937 13.5 16.5 13.5C13.7386 13.5 11.5 11.2614 11.5 8.5C11.5 6.8599 12.2892 5.40423 13.5106 4.49167C13.6721 4.37101 13.7453 4.16516 13.6963 3.96961C13.6473 3.77406 13.4857 3.62706 13.2864 3.59678C12.8666 3.53302 12.437 3.5 12 3.5Z"
        fill="currentColor"
      /></svg
    >
  </button>
  <button
    class="flex {currentTheme === 'light'
      ? ''
      : 'hidden'} items-center text-slate-800 dark:text-white"
    formaction="/?/setTheme&theme=dark"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="1.5"
      stroke="currentColor"
      class="w-5 h-5"
      ><path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M12 3v2.25m6.364.386l-1.591 1.591M21 12h-2.25m-.386 6.364l-1.591-1.591M12 18.75V21m-4.773-4.227l-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z"
      /></svg
    >
  </button>
</form>
