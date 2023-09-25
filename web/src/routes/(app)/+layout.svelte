<script lang="ts">
  import Nav from '$lib/components/Nav.svelte';
</script>

<div class="container mx-auto">
  <Nav />
  <main class="px-6 py-44 mx-auto max-w-[700px]">
    <slot />
  </main>
</div>
