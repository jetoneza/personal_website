<script lang="ts">
	import { page } from '$app/stores';
</script>

<nav class="main-nav flex pt-6 pb-4 px-10 justify-between border-b">
	<a href="/" class="text-2xl font-serif"
		><span class="font-bold hover:text-cyan-600">Jet</span> Ordaneza</a
	>
	<ul class="flex space-x-10 font-mono justify-end">
		<!-- TODO: Show links if content is available -->
		<!-- <li aria-current={$page.url.pathname === '/blog' ? 'page' : undefined}> -->
		<!-- 	<a href="/blog">Blog</a> -->
		<!-- </li> -->
		<!-- <li aria-current={$page.url.pathname === '/about' ? 'page' : undefined}> -->
		<!-- 	<a href="/about">About</a> -->
		<!-- </li> -->
		<li aria-current={$page.url.pathname === '/contact' ? 'page' : undefined}>
			<a href="/contact">Contact</a>
		</li>
		<li>
			<a href="https://github.com/jetoneza" target="_blank">Github</a>
		</li>
	</ul>
</nav>

<style lang="postcss">
	li {
		@apply hover:font-bold hover:text-cyan-600;
	}

	li[aria-current='page'] {
		@apply font-bold text-cyan-600;
	}

	.main-nav {
		view-transition-name: header;
	}
</style>
