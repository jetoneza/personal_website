<script lang="ts">
  import { onNavigate } from '$app/navigation';

  // Styles
  import '@fontsource/source-sans-pro/400.css';
  import '@fontsource/source-sans-pro/700.css';
  import '@fontsource/source-sans-pro/900.css';
  import '@fontsource-variable/open-sans';
  import '../app.postcss';

  onNavigate((navigation) => {
    if (!document.startViewTransition) return;

    return new Promise((resolve) => {
      document.startViewTransition(async () => {
        resolve();
        await navigation.complete;
      });
    });
  });
</script>

<slot />
