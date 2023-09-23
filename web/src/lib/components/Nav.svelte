<script lang="ts">
  import { afterNavigate } from '$app/navigation';
  import { page } from '$app/stores';
  import HamburgerMenu from './HamburgerMenu.svelte';

  let open = false;

  afterNavigate(() => {
    open = false;
  });

  function toggleMenu() {
    open = !open;
  }
</script>

<nav class="main-nav fixed top-0 right-0 left-0 border-b bg-white">
  <div class="container mx-auto flex p-4 sm:pt-6 sm:px-10 justify-between">
    <a href="/" class="text-2xl font-sans-pro">
      <span class="font-bold hover:text-cyan-600">Jet</span> Ordaneza
    </a>

    <HamburgerMenu {open} onClick={toggleMenu} />

    <ul
      class="{open ? 'flex' : 'hidden'} absolute bg-white top-20 left-0 right-0
    bottom-0 px-4 space-y-6 flex-col sm:p-0 sm:top-0 sm:space-y-0 sm:relative
    sm:flex sm:flex-row sm:space-x-10 sm:justify-end items-center"
    >
      <!-- TODO: Show links if content is available -->
      <li
        class="text-lg sm:text-base"
        aria-current={$page.url.pathname === '/blog' ? 'page' : undefined}
      >
        <a href="/blog">Blog</a>
      </li>
      <li
        class="text-lg sm:text-base"
        aria-current={$page.url.pathname === '/about' ? 'page' : undefined}
      >
        <a href="/about">About</a>
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
