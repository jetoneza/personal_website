<script lang="ts">
  import { goto } from '$app/navigation';

  let email = '';
  let password = '';
  let error = '';

  async function handleSubmit() {
    error = '';

    const data = {
      email,
      password,
    };

    const response = await fetch('/api/v1/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    const result = JSON.parse(await response.text());

    if (result.status === 'fail') {
      error = result.error;
      return;
    }

    goto('/admin/dashboard');
  }
</script>

<div class="container login mx-auto p-6 flex flex-col justify-center items-center h-screen">
  <div class="form-wrapper w-80 p-6 rounded-lg drop-shadow-lg bg-white font-mono">
    <div class="title font-bold text-lg">Login to Dashboard</div>

    {#if error}
      <p class="p-4 bg-pink-100 border border-red-500 rounded-lg text-red-500 my-2 text-sm">
        {error}
      </p>
    {/if}

    <form class="flex flex-col space-y-4 mt-4" on:submit|preventDefault={handleSubmit}>
      <div class="input-wrapper">
        <label for="email" class="block mb-2 text-sm font-medium dark:text-white"
          >Email address</label
        >
        <input
          type="email"
          id="email"
          name="email"
          bind:value={email}
          class="
            border border-gray-300 rounded-lg text-sm block w-full p-2.5
            focus:ring-2 focus:outline-none focus:ring-blue-300
          "
          required
        />
      </div>
      <div class="input-wrapper">
        <label for="password" class="block mb-2 text-sm font-medium dark:text-white">
          Password
        </label>
        <input
          type="password"
          id="password"
          name="password"
          bind:value={password}
          class="
            border border-gray-300 rounded-lg text-sm block w-full p-2.5
            focus:ring-2 focus:outline-none focus:ring-blue-300
          "
          required
        />
      </div>
      <button
        type="submit"
        class="
          text-white bg-cyan-600 hover:bg-cyan-700 focus:ring-4
          focus:outline-none focus:ring-blue-300 font-medium rounded-lg
          text-sm w-full sm:w-auto px-5 py-2.5 text-center
        "
      >
        Submit
      </button>
    </form>
  </div>
</div>
