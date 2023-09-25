<nav
  class="main-nav fixed flex top-0 right-0 left-0 justify-between border-b p-6 dark:bg-zinc-800 dark:border-b-zinc-700"
>
  <a href="/admin/dashboard" class="text-2xl"
    ><span class="font-bold hover:text-cyan-600">Admin</span> Dashboard</a
  >

  <ul class="flex gap-4 items-center">
    <li class="text-lg sm:text-base">
      <a href="/">Home</a>
    </li>
    <li class="text-lg sm:text-base">
      <a href="/admin/logout">Logout</a>
    </li>
  </ul>
</nav>

<style lang="postcss">
  li {
    @apply hover:text-cyan-600;
  }

  .main-nav {
    view-transition-name: header;
  }
</style>
