<script lang="ts">
  import { page } from '$app/stores';

  import CalendarMonthSolid from 'flowbite-svelte-icons/CalendarMonthSolid.svelte';
  import HomeSolid from 'flowbite-svelte-icons/HomeSolid.svelte';
  import PenNibSolid from 'flowbite-svelte-icons/PenNibSolid.svelte';

  import ThemeSwitch from '../ThemeSwitch.svelte';
</script>

<aside
  class="side-nav hidden md:fixed bg-zinc-100 top-0 left-0 z-40 w-64 h-screen border-r dark:border-r-zinc-800 dark:bg-zinc-900 md:flex flex-col"
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
          class="flex items-center px-4 py-2 rounded-lg hover:bg-gray-100 dark:hover:bg-zinc-600 group gap-3"
          aria-current={$page.url.pathname === '/admin/dashboard' ? 'page' : undefined}
        >
          <HomeSolid size="md" />
          <span>Home</span>
        </a>
      </li>
      <li>
        <a
          href="/admin/dashboard/blog"
          class="flex items-center px-4 py-2 rounded-lg hover:bg-gray-100 dark:hover:bg-zinc-600 group gap-3"
          aria-current={$page.url.pathname === '/admin/dashboard/blog' ? 'page' : undefined}
        >
          <PenNibSolid size="md" />
          <span>Blog</span>
        </a>
      </li>
      <li>
        <a
          href="/admin/dashboard/calendar"
          class="flex items-center px-4 py-2 rounded-lg hover:bg-gray-100 dark:hover:bg-zinc-600 group gap-3"
          aria-current={$page.url.pathname.includes('/admin/dashboard/calendar')
            ? 'page'
            : undefined}
        >
          <CalendarMonthSolid size="md" />
          <span>Calendar</span>
        </a>
      </li>
    </ul>
  </div>

  <div class="actions flex mt-auto w-full justify-center text-center py-6 px-4 gap-4">
    <div class="bg-zinc-200 p-2 rounded dark:bg-zinc-600">
      <ThemeSwitch />
    </div>
    <a href="/admin/logout" class="bg-zinc-200 p-2 rounded flex items-center dark:bg-zinc-600">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        height="1em"
        viewBox="0 0 512 512"
        class="fill-gray-900 dark:fill-white"
      >
        <path
          d="M497 273L329 441c-15 15-41 4.5-41-17v-96H152c-13.3
          0-24-10.7-24-24v-96c0-13.3 10.7-24 24-24h136V88c0-21.4
          25.9-32 41-17l168 168c9.3 9.4 9.3 24.6 0 34zM192
          436v-40c0-6.6-5.4-12-12-12H96c-17.7 0-32-14.3-32-32V160c0-17.7
          14.3-32 32-32h84c6.6 0 12-5.4 12-12V76c0-6.6-5.4-12-12-12H96c-53
          0-96 43-96 96v192c0 53 43 96 96 96h84c6.6 0 12-5.4 12-12z"
        />
      </svg>
    </a>
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
