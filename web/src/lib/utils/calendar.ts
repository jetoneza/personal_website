// TODO: declare correct specific fields
type EventTypes = {
  [key: string]: {
    [key: string]: string;
  };
};

const types: EventTypes = {
  work: {
    backgroundColor: '#0891b2',
  },
  task: {
    backgroundColor: '#14b8a6',
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
  if (!type) {
    return {};
  }

  return types[type];
};
