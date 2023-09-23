<script lang="ts">
  import Nav from '$lib/components/Nav.svelte';
</script>

<div class="container mx-auto bg-white">
  <Nav />
  <div class="px-4 py-28">
    <slot />
  </div>
</div>
