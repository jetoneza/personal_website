<script lang="ts">
  import SideNav from '$lib/components/admin/SideNav.svelte';
  import Nav from '$lib/components/admin/Nav.svelte';
</script>

<div class="dashboard">
  <Nav />
  <div class="relative top-[81px]">
    <SideNav />

    <div class="p-6 sm:ml-64">
      <slot />
    </div>
  </div>
</div>
