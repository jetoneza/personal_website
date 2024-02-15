<script lang="ts">
  import Alerts from '$lib/components/admin/Alerts.svelte';
  import SideNav from '$lib/components/admin/SideNav.svelte';
</script>

<svelte:head>
  <title>Dashboard</title>
</svelte:head>

<div class="dashboard relative">
  <SideNav />

  <div class="px-6 py-8 sm:ml-64">
    <slot />
  </div>
</div>

<Alerts />
