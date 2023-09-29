<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<svelte:head>
  <title>{data.meta.title}</title>

  <meta name="description" content={data.meta.description} />
  <meta name="keywords" content={data.meta.keywords} />

  <!-- Social media -->
  <meta name="twitter:card" content="summary" />
  <meta name="twitter:title" content={data.meta.title} />
  <meta name="twitter:description" content={data.meta.description} />
  <meta name="twitter:image" content={data.meta.imageUrl} />
  <meta name="twitter:url" content={data.meta.url} />

  <meta property="og:title" content={data.meta.title} />
  <meta property="og:description" content={data.meta.description} />
  <meta property="og:image" content={data.meta.imageUrl} />
  <meta property="og:url" content={data.meta.url} />
  <meta property="og:type" content="website" />
</svelte:head>

<div class="flex flex-col gap-16 md:gap-8">
  <div class="flex flex-col gap-2">
    <h1 class="font-bold text-4xl font-sans-pro">Blog</h1>
    <p class="text-secondary animate-in">What I think about stuff.</p>
  </div>
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
                    <a href="/blog/{post.slug}">{post.title}</a>
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
</div>
