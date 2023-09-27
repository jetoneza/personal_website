<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<div class="header space-y-8">
  <h1 class="name font-bold text-4xl font-sans-pro">
    Hello, I'm Jet<span class="text-cyan-600">.</span>
  </h1>

  <div class="content space-y-4">
    <p class="text-justify">
      A <span class="font-bold">Software Engineer</span>. I love building cool things with code. I'm
      always curious about how things work and love exploring new stuff.
    </p>

    <p class="text-justify">
      I'm excited to share my journey with you, whether it's my coding adventures or my interests
      outside of work. Let's explore the world together!
    </p>
  </div>
</div>

<div class="posts mt-16">
  <h1 class="font-bold text-2xl font-sans-pro">Latest posts</h1>
  <ul class="space-y-4 mt-8">
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
  <div class="mt-8">
    <a href="/blog" class="font-bold underline">See more</a>
  </div>
</div>
