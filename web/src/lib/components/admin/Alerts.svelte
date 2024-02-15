<script lang="ts">
  import { Alert } from 'flowbite-svelte';

  import { addAlert, alertsStore, type AlertType } from '$lib/stores/alerts';

  type Color =
    | 'green'
    | 'red'
    | 'blue'
    | 'form'
    | 'none'
    | 'gray'
    | 'yellow'
    | 'indigo'
    | 'purple'
    | 'pink'
    | 'light'
    | 'dark'
    | 'default'
    | 'dropdown'
    | 'navbar'
    | 'navbarUl'
    | 'primary'
    | 'orange'
    | undefined;

  let alerts: AlertType[] = [];

  alertsStore.subscribe((a) => {
    alerts = a;
  });

  function getAlertColor(type: string) {
    const colors: { [key: string]: string } = {
      success: 'green',
      error: 'red',
      info: 'blue',
    };

    return (colors[type] || colors.info) as Color;
  }

  function getTitle(type: string) {
    const title: { [key: string]: string } = {
      success: 'Success!',
      error: 'Error!',
      info: 'Info!',
    };

    return title[type] || title.info;
  }
</script>

<div class="alerts-group fixed right-10 bottom-10 flex flex-col gap-2">
  {#each alerts as alert}
    <Alert color={getAlertColor(alert.type)}>
      <span class="font-semibold">{getTitle(alert.type)}</span>
      {alert.message}
    </Alert>
  {/each}
</div>
