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

    <HamburgerMenu {open} onClick={toggleMenu} />

    <ul
      class="{open ? 'flex' : 'hidden'} absolute top-20 left-0 right-0
    bottom-0 px-4 space-y-6 flex-col sm:p-0 sm:top-0 sm:space-y-0 sm:relative
    sm:flex sm:flex-row sm:space-x-10 sm:justify-end items-center"
    >
      <li
        class="text-lg sm:text-base"
        aria-current={$page.url.pathname === '/about' ? 'page' : undefined}
      >
        <a href="/about">About</a>
      </li>
      <li
        class="text-lg sm:text-base"
        aria-current={$page.url.pathname.includes('/blog') ? 'page' : undefined}
      >
        <a href="/blog">Blog</a>
      </li>
      <li
        class="text-lg sm:text-base"
        aria-current={$page.url.pathname === '/contact' ? 'page' : undefined}
      >
        <a href="/contact">Contact</a>
      </li>
      <li class="text-lg sm:text-base">
        <a href="https://github.com/jetoneza" target="_blank">Github</a>
      </li>
      <li class="hover:bg-zinc-100 hover:dark:bg-zinc-700 p-2 rounded-lg">
        <ThemeSwitch />
      </li>
    </ul>
  </div>
</nav>

<style lang="postcss">
  li {
    @apply hover:text-cyan-600;
  }

  li[aria-current='page'] {
    @apply text-cyan-600;
  }

  .main-nav {
    view-transition-name: header;
  }
</style>
