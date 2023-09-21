<script lang="ts">
	const links = [
		{
			url: '/dashboard',
			title: 'Dashboard',
		},
		{
			url: '/blog',
			title: 'Blog',
		},
	];
</script>

<aside
	class="absolute font-mono top-0 left-0 z-40 w-64 h-screen"
>
	<div class="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
		<ul class="space-y-2 font-medium">
			{#each links as link}
				<li />
				<li>
					<a
						href={link.url}
						class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
					>
						<span class="ml-3">{link.title}</span>
					</a>
				</li>
			{/each}
		</ul>
	</div>
</aside>
