import { writable, get } from 'svelte/store';

export type AlertType = {
  id?: number;
  type: string;
  message: string;
};

const DISPLAY_TIME = 5e3; // 5 seconds

export const alertsStore = writable<AlertType[]>([]);

export function addAlert(alert: AlertType) {
  const id = Date.now();

  alertsStore.update((alerts) => [
    ...alerts,
    {
      ...alert,
      id,
    },
  ]);

  setTimeout(() => {
    const alerts = get(alertsStore) || [];
    const filteredAlerts = alerts.filter((a) => a.id !== id);

    alertsStore.update(() => filteredAlerts);
  }, DISPLAY_TIME);
}
