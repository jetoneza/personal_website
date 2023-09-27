<script lang="ts">
  import { formatDate } from '$lib/utils/date';
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<main class="blog-dashboard">
  <div class="actions w-full flex justify-between items-center">
    <h1 class="font-bold text-2xl font-sans-pro">Posts</h1>
    <a class="btn text-sm" href="/admin/dashboard/blog/create">Create Post</a>
  </div>

  <div class="relative overflow-x-auto mt-6">
    <table class="w-full text-sm text-left text-zinc-500 dark:text-zinc-400">
      <thead class="text-sm text-zinc-700 border-b dark:border-zinc-700 dark:text-zinc-400">
        <tr>
          <th scope="col" class="px-6 py-3"> Title </th>
          <th scope="col" class="px-6 py-3"> Published </th>
          <th scope="col" class="px-6 py-3"> Date </th>
        </tr>
      </thead>
      <tbody>
        {#each data.posts as post}
          <tr
            class="bg-white border-b dark:bg-zinc-800 dark:border-zinc-700 hover:bg-zinc-100 dark:hover:bg-zinc-700"
          >
            <th
              scope="row"
              class="px-6 py-6 font-medium text-zinc-900 whitespace-nowrap dark:text-white"
            >
              <a href="/admin/dashboard/blog/{post.id}/edit">{post.title}</a>
            </th>
            <td class="px-6 py-4">
              <span
                class="badge {post.published
                  ? 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300'
                  : 'bg-zinc-100 dark:bg-zinc-700 dark:text-zinc-300'}"
              >
                {post.published ? 'Published' : 'Unpublished'}
              </span>
            </td>
            <td class="px-6 py-4">
              {formatDate(post.createdAt)}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</main>
