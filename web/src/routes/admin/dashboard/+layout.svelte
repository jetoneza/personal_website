<script lang="ts">
  import SideNav from '$lib/components/admin/SideNav.svelte';
</script>

<div class="dashboard relative">
  <SideNav />

  <div class="px-6 py-8 sm:ml-64">
    <slot />
  </div>
</div>
