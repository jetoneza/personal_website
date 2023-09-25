<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<article class="prose dark:prose-invert">
  {@html data.post.content}
</article>
