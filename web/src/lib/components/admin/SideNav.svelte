<script lang="ts">
  import { page } from '$app/stores';

  const links = [
    {
      url: '/admin/dashboard',
      title: 'Home',
    },
    {
      url: '/admin/dashboard/blog',
      title: 'Blog',
    },
  ];
</script>

<aside class="side-nav absolute top-0 left-0 z-40 w-64 h-screen border-r dark:border-r-zinc-700">
  <div class="h-full px-3 py-4 overflow-y-auto dark:bg-zinc-800">
    <ul class="space-y-2 font-medium">
      {#each links as link}
        <li />
        <li>
          <a
            href={link.url}
            class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-zinc-600 group"
            aria-current={$page.url.pathname === link.url ? 'page' : undefined}
          >
            <span class="ml-3">{link.title}</span>
          </a>
        </li>
      {/each}
    </ul>
  </div>
</aside>

<style lang="postcss">
  a[aria-current='page'] {
    @apply bg-gray-100 dark:bg-zinc-600;
  }

  .side-nav {
    view-transition-name: sidenav;
  }
</style>
