<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<div class="flex flex-col gap-16 md:gap-24">
  <div class="flex flex-col gap-2">
    <h1 class="animate-in text-3xl font-bold tracking-tight">Blog</h1>
    <p class="text-secondary animate-in">What I think about stuff.</p>
  </div>
  <ul class="space-y-4">
    {#each data.posts as post}
      <li class="flex gap-10">
        <div class="italic">
          <time datetime={post.createdAt}>{post.formattedCreatedAt}</time>
        </div>
        <!-- TODO: Use slug for links -->
        <a href="/blog/{post.slug}" class="font-bold">
          {post.title} <br />
        </a>
      </li>
    {/each}
  </ul>
</div>
