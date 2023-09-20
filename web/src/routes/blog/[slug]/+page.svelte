<script lang="ts">
	import type { PageData } from '../../$types';
	import Content from './index.md';

	export let data: PageData;

	console.log(data);
</script>

<article class="prose">
	<Content />
</article>
