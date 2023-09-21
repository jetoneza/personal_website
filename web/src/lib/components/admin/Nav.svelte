<nav class="admin-nav flex p-4 sm:pt-6 sm:px-10 justify-between border-b">
	<a href="/admin/dashboard" class="text-2xl font-serif"
		><span class="font-bold hover:text-cyan-600">Admin</span> Dashboard</a
	>

	<ul
		class="flex absolute bg-white top-20 left-0 right-0
    bottom-0 px-4 space-y-6 flex-col sm:p-0 sm:top-0 sm:space-y-0 sm:relative
    sm:flex sm:flex-row sm:space-x-10 font-mono sm:justify-end items-center"
	>
		<li class="text-lg sm:text-base">
			<a href="/">Home</a>
		</li>
		<li class="text-lg sm:text-base">
			<a href="/logout">Logout</a>
		</li>
	</ul>
</nav>

<style lang="postcss">
	li {
		@apply hover:font-bold hover:text-cyan-600;
	}

	.admin-nav {
		view-transition-name: header;
	}
</style>
