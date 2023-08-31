<script>
	import '../app.postcss';
</script>

<nav class="py-6 px-10">
	<ul class="flex space-x-10 font-mono justify-end">
		<li><a href="/">Home</a></li>
		<li><a href="/blog">Blog</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
	</ul>
</nav>

<slot />

<style lang="postcss">
	li {
		@apply hover:font-bold;
	}
</style>
