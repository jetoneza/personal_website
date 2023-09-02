<script lang="ts">
	export let open = false;
	export let onClick: () => void;
</script>

<button class="relative w-8 sm:hidden" on:click={onClick}>
	<span
		class="block absolute w-8 h-0.5 bg-slate-800 transform transition duration-500 ease-in-out {open
			? 'rotate-45'
			: ''} {!open ? '-translate-y-1.5' : ''}"
	/>
	<span
		class="block absolute w-8 h-0.5 bg-slate-800 transform transition duration-500 ease-in-out {open
			? 'opacity-0'
			: ''}"
	/>
	<span
		class="block absolute w-8 h-0.5 bg-slate-800 transform transition duration-500 ease-in-out {open
			? '-rotate-45'
			: ''} {!open ? 'translate-y-1.5' : ''}"
	/>
</button>
