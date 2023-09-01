<script lang="ts">
	import { onNavigate } from '$app/navigation';
	import Nav from '$lib/components/Nav.svelte';

	// Styles
	import '../app.postcss';

	onNavigate((navigation) => {
		if (!document.startViewTransition) return;

		return new Promise((resolve) => {
			document.startViewTransition(async () => {
				resolve();
				await navigation.complete;
			});
		});
	});
</script>

<Nav />

<div class="container mx-auto px-4 py-10">
	<slot />
</div>
