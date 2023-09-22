<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<ul>
  {#each data.posts as post}
    <li>
      <!-- TODO: Use slug for links -->
      <a href="/blog/{post.id}" class="font-bold font-serif text-2xl">
        {post.title} <br />
        <span class="text-gray-700 text-xs">{post.createdAt}</span>
      </a>
    </li>
  {/each}
</ul>
