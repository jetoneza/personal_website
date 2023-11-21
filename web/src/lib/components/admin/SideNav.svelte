<script lang="ts">
  import { page } from '$app/stores';

  import ThemeSwitch from '../ThemeSwitch.svelte';
</script>

<aside
  class="side-nav fixed bg-zinc-100 top-0 left-0 z-40 w-64 h-screen border-r dark:border-r-zinc-800 dark:bg-zinc-900 flex flex-col"
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
          class="flex items-center p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-zinc-600 group gap-2"
          aria-current={$page.url.pathname === '/admin/dashboard' ? 'page' : undefined}
        >
          <span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              height="1em"
              width="1em"
              viewBox="0 0 576 512"
              class="fill-gray-900 dark:fill-white"
            >
              <path
                d="M575.8 255.5c0 18-15 32.1-32 32.1h-32l.7 160.2c0 2.7-.2 5.4-.5
              8.1V472c0 22.1-17.9 40-40 40H456c-1.1 0-2.2 0-3.3-.1c-1.4 .1-2.8
              .1-4.2 .1H416 392c-22.1 0-40-17.9-40-40V448 384c0-17.7-14.3-32-32-32H256c-17.7
              0-32 14.3-32 32v64 24c0 22.1-17.9 40-40 40H160 128.1c-1.5 0-3-.1-4.5-.2c-1.2
              .1-2.4 .2-3.6 .2H104c-22.1 0-40-17.9-40-40V360c0-.9 0-1.9
              .1-2.8V287.6H32c-18 0-32-14-32-32.1c0-9 3-17 10-24L266.4 8c7-7
              15-8 22-8s15 2 21 7L564.8 231.5c8 7 12 15 11 24z"
              />
            </svg>
          </span>
          <span>Home</span>
        </a>
      </li>
      <li>
        <a
          href="/admin/dashboard/blog"
          class="flex items-center p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-zinc-600 group gap-2"
          aria-current={$page.url.pathname === '/admin/dashboard/blog' ? 'page' : undefined}
        >
          <span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              height="1em"
              width="1em"
              viewBox="0 0 512 512"
              class="fill-gray-900 dark:fill-white"
            >
              <path
                d="M471.6 21.7c-21.9-21.9-57.3-21.9-79.2 0L362.3 51.7l97.9 97.9
                30.1-30.1c21.9-21.9 21.9-57.3 0-79.2L471.6 21.7zm-299.2 220c-6.1
                6.1-10.8 13.6-13.5 21.9l-29.6 88.8c-2.9 8.6-.6 18.1 5.8 24.6s15.9
                8.7 24.6 5.8l88.8-29.6c8.2-2.7 15.7-7.4 21.9-13.5L437.7 172.3
                339.7 74.3 172.4 241.7zM96 64C43 64 0 107 0 160V416c0 53 43 96
                96 96H352c53 0 96-43 96-96V320c0-17.7-14.3-32-32-32s-32 14.3-32
                32v96c0 17.7-14.3 32-32 32H96c-17.7 0-32-14.3-32-32V160c0-17.7
                14.3-32 32-32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32H96z"
              />
            </svg>
          </span>
          <span>Blog</span>
        </a>
      </li>
      <li>
        <a
          href="/admin/dashboard/calendar"
          class="flex items-center p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-zinc-600 group gap-2"
          aria-current={$page.url.pathname === '/admin/dashboard/calendar' ? 'page' : undefined}
        >
          <span>
            <svg
              class="text-gray-800 dark:text-white"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              height="1em"
              width="1em"
              viewBox="0 0 20 20"
            >
              <path
                fill="currentColor"
                d="M6 1a1 1 0 0 0-2 0h2ZM4 4a1 1 0 0 0 2 0H4Zm7-3a1 1 0 1 0-2 0h2ZM9 4a1 1 0 1 0 2 0H9Zm7-3a1 1 0 1 0-2 0h2Zm-2 3a1 1 0 1 0 2 0h-2ZM1 6a1 1 0 0 0 0 2V6Zm18 2a1 1 0 1 0 0-2v2ZM5 11v-1H4v1h1Zm0 .01H4v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM10 11v-1H9v1h1Zm0 .01H9v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM10 15v-1H9v1h1Zm0 .01H9v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM15 15v-1h-1v1h1Zm0 .01h-1v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM15 11v-1h-1v1h1Zm0 .01h-1v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM5 15v-1H4v1h1Zm0 .01H4v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM2 4h16V2H2v2Zm16 0h2a2 2 0 0 0-2-2v2Zm0 0v14h2V4h-2Zm0 14v2a2 2 0 0 0 2-2h-2Zm0 0H2v2h16v-2ZM2 18H0a2 2 0 0 0 2 2v-2Zm0 0V4H0v14h2ZM2 4V2a2 2 0 0 0-2 2h2Zm2-3v3h2V1H4Zm5 0v3h2V1H9Zm5 0v3h2V1h-2ZM1 8h18V6H1v2Zm3 3v.01h2V11H4Zm1 1.01h.01v-2H5v2Zm1.01-1V11h-2v.01h2Zm-1-1.01H5v2h.01v-2ZM9 11v.01h2V11H9Zm1 1.01h.01v-2H10v2Zm1.01-1V11h-2v.01h2Zm-1-1.01H10v2h.01v-2ZM9 15v.01h2V15H9Zm1 1.01h.01v-2H10v2Zm1.01-1V15h-2v.01h2Zm-1-1.01H10v2h.01v-2ZM14 15v.01h2V15h-2Zm1 1.01h.01v-2H15v2Zm1.01-1V15h-2v.01h2Zm-1-1.01H15v2h.01v-2ZM14 11v.01h2V11h-2Zm1 1.01h.01v-2H15v2Zm1.01-1V11h-2v.01h2Zm-1-1.01H15v2h.01v-2ZM4 15v.01h2V15H4Zm1 1.01h.01v-2H5v2Zm1.01-1V15h-2v.01h2Zm-1-1.01H5v2h.01v-2Z"
              />
            </svg>
          </span>
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
