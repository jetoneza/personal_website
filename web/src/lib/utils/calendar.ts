// TODO: declare correct specific fields
type EventTypes = {
  [key: string]: {
    [key: string]: string;
  };
};

const types: EventTypes = {
  work: {
    backgroundColor: '#0891b2',
    badgeClass: 'bg-cyan-100 text-cyan-800 dark:bg-cyan-900 dark:text-cyan-300',
  },
  task: {
    backgroundColor: '#14b8a6',
    badgeClass: 'bg-teal-100 text-teal-800 dark:bg-teal-900 dark:text-teal-300',
  },
  'regular-holiday': {
    backgroundColor: '#22c55e',
    badgeClass: 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300',
  },
  'special-holiday': {
    backgroundColor: '#fbbf24',
    badgeClass: 'bg-amber-100 text-amber-800 dark:bg-amber-900 dark:text-amber-300',
  },
};

export const calendarOptions = {
  view: 'dayGridMonth',
  headerToolbar: {
    start: 'prev,next today',
    center: 'title',
    end: 'dayGridMonth,timeGridWeek,timeGridDay',
  },
  buttonText: {
    today: 'Today',
    dayGridMonth: 'Month',
    timeGridDay: 'Day',
    timeGridWeek: 'Week',
  },
  scrollTime: '09:00:00',
  views: {
    timeGridWeek: { pointer: true },
    resourceTimeGridWeek: { pointer: true },
  },
  nowIndicator: true,
  selectable: true,
  height: '75%',
  display: 'background',
};

export const getEventColor = (type?: string) => {
  const defaultColor = {
    backgroundColor: '#3b82f6',
    badgeClass: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-300',
  };

  if (!type) {
    return defaultColor;
  }

  return types[type] ?? defaultColor;
};
