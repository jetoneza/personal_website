<script lang="ts">
  import SideNav from '$lib/components/admin/SideNav.svelte';
  import Nav from '$lib/components/admin/Nav.svelte';
</script>

<Nav />
<div class="relative">
  <SideNav />

  <div class="p-4 sm:ml-64">
    <slot />
  </div>
</div>
