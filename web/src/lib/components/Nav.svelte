<script lang="ts">
  import { afterNavigate } from '$app/navigation';
  import { page } from '$app/stores';
  import HamburgerMenu from './HamburgerMenu.svelte';
  import ThemeSwitch from './ThemeSwitch.svelte';

  let open = false;

  afterNavigate(() => {
    open = false;
  });

  const toggleMenu = () => {
    open = !open;
  };
</script>

<nav
  class="main-nav fixed top-0 right-0 left-0 border-b bg-white dark:bg-zinc-800 dark:border-b-zinc-800 dark:drop-shadow"
>
  <div class="container mx-auto flex p-4 sm:pt-6 justify-between max-w-[700px]">
    <a href="/" class="text-2xl font-sans-pro">
      <span class="font-bold hover:text-cyan-600">Jet</span> Ordaneza
    </a>

    <div class="flex space-x-4 sm:hidden">
      <HamburgerMenu {open} onClick={toggleMenu} />
      <div class="border border-zinc-200 dark:border-zinc-700 p-2 rounded-lg">
        <ThemeSwitch />
      </div>
    </div>

    <ul
      class="{open ? 'block' : 'hidden'} absolute top-[78px] right-[60px]
      bg-white rounded-lg w-1/2 p-2 py-4 space-y-2 flex-col drop-shadow-lg border
      dark:bg-zinc-800 dark:border-zinc-800 sm:border-0 sm:drop-shadow-none
      sm:p-0 sm:top-0 sm:space-y-0 sm:relative sm:flex sm:flex-row sm:space-x-10 sm:justify-end items-center"
    >
      <li
        class="sm:hidden px-4 py-2 sm:p-0 sm:font-semibold"
        aria-current={$page.url.pathname === '/' ? 'page' : undefined}
      >
        <a href="/">Home</a>
      </li>
      <li
        class="px-4 py-2 sm:p-0 sm:font-semibold"
        aria-current={$page.url.pathname === '/about' ? 'page' : undefined}
      >
        <a href="/about">About</a>
      </li>
      <li
        class="px-4 py-2 sm:p-0 sm:font-semibold"
        aria-current={$page.url.pathname.includes('/blog') ? 'page' : undefined}
      >
        <a href="/blog">Blog</a>
      </li>
      <li
        class="px-4 py-2 sm:p-0 sm:font-semibold"
        aria-current={$page.url.pathname === '/contact' ? 'page' : undefined}
      >
        <a href="/contact">Contact</a>
      </li>
      <li class="px-4 py-2 sm:p-0 sm:font-semibold">
        <a href="https://github.com/jetoneza" target="_blank">Github</a>
      </li>
      <li class="hidden sm:block hover:bg-zinc-100 hover:dark:bg-zinc-700 sm:p-2 rounded-lg">
        <ThemeSwitch />
      </li>
    </ul>
  </div>
</nav>

<style lang="postcss">
  li {
    @apply sm:hover:text-cyan-600;
  }

  li[aria-current='page'] {
    @apply bg-gray-200 dark:bg-zinc-700 sm:bg-transparent sm:dark:bg-transparent rounded-lg sm:rounded-none sm:text-cyan-500;
  }

  .main-nav {
    view-transition-name: header;
  }
</style>
