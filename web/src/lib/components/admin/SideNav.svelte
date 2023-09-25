<script lang="ts">
  import { page } from '$app/stores';
</script>

<aside
  class="side-nav absolute bg-zinc-100 top-0 left-0 z-40 w-64 h-screen border-r dark:border-r-zinc-800 dark:bg-zinc-900 flex flex-col"
>
  <div class="header w-full text-center py-6 px-4">
    <a href="/admin/dashboard" class="text-2xl"
      ><span class="font-bold hover:text-cyan-600 font-sans-pro">Admin</span> Dashboard</a
    >
  </div>
  <div class="px-3 py-4 overflow-y-auto">
    <ul class="space-y-2 font-medium">
      <li>
        <a
          href="/admin/dashboard"
          class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-zinc-600 group"
          aria-current={$page.url.pathname === '/admin/dashboard' ? 'page' : undefined}
        >
          <span class="ml-3">Home</span>
        </a>
      </li>
      <li>
        <a
          href="/admin/dashboard/blog"
          class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-zinc-600 group"
          aria-current={$page.url.pathname === '/admin/dashboard/blog' ? 'page' : undefined}
        >
          <span class="ml-3">Blog</span>
        </a>
      </li>
    </ul>
  </div>

  <div class="actions mt-auto w-full text-center py-6 px-4">
    <a href="/admin/logout">Logout</a>
  </div>
</aside>

<style lang="postcss">
  a[aria-current='page'] {
    @apply bg-gray-200 dark:bg-zinc-600;
  }

  .side-nav {
    view-transition-name: sidenav;
  }
</style>
