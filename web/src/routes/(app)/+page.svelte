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
  <ul class="divide-y divide-gray-200 dark:divide-gray-700">
    {#each data.posts as post}
      <li class="py-8">
        <article>
          <div class="space-y-2">
            <div class="text-base font-medium leading-6 text-gray-500 dark:text-gray-400">
              <time datetime={post.createdAt}>{post.formattedCreatedAt}</time>
            </div>
            <div class="space-y-3">
              <div class="space-y-4">
                <div>
                  <h2 class="text-xl font-bold font-sans-pro leading-8 tracking-tight">
                    <a
                      href="/blog/{post.slug}"
                      >{post.title}</a
                    >
                  </h2>
                </div>
                <div class="prose max-w-none dark:prose-invert">
                  {post.description ?? ''}
                </div>
              </div>
              <div class="text-base font-medium leading-6">
                <a
                  class="text-cyan-500 hover:text-cyan-600 dark:hover:text-cyan-400"
                  aria-label={post.title}
                  href="/blog/{post.slug}">Read more →</a
                >
              </div>
            </div>
          </div>
        </article>
      </li>
    {/each}
  </ul>
  <div class="mt-8">
    <a href="/blog" class="font-bold underline">See more</a>
  </div>
</div>
