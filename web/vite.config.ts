import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

export default defineConfig({
  plugins: [sveltekit()],
  test: {
    include: ['src/**/*.{test,spec}.{js,ts}'],
  },
  clearScreen: false,
  build: {
    sourcemap: false,
  },
  server: {
    proxy: {
      // TODO: Use env. This proxy is only for development.
      '/api': 'http://127.0.0.1:3000',
    },
  },
});
