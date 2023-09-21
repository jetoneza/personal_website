<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<article class="prose">
  {@html data.post.content}
</article>
