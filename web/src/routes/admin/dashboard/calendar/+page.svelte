<script lang="ts">
  // Libraries

  // TODO: Find workaround on sveltekit not reading declared modules in app.d.ts within .svelte files
  // @ts-expect-error: No declaration file for module
  import Calendar from '@event-calendar/core';
  // @ts-expect-error: No declaration file for module
  import TimeGrid from '@event-calendar/time-grid';
  // @ts-expect-error: No declaration file for module
  import DayGrid from '@event-calendar/day-grid';

  // Styles
  import './styles.css';

  // TODO: declare correct specific fields
  type Event = {
    [key: string]: string | number | undefined;
  };

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

  const events = [
    {
      id: 3,
      start: new Date('12-25-2023'),
      end: new Date('12-25-2023'),
      allDay: true,
      notes: 'Christmas 🌲🎁🎉',
      type: 'holiday',
      createdAt: new Date(),
    },
  ];

  const getEventColor = (type?: string) => {
    if (!type) {
      return {};
    }

    return types[type];
  };

  const handleEventClick = (info: { event: Event }) => {
    // TODO: Open event modal
    // TODO: Display event
    console.log(info.event.id);
  };

  const plugins = [TimeGrid, DayGrid];
  const options = {
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
    eventClick: handleEventClick,
    events: events.map(({ id, start, end, allDay, notes, type }) => ({
      id,
      start,
      end,
      allDay,
      title: notes,
      ...getEventColor(type),
    })),
  };
</script>

<div class="work h-screen">
  <h1 class="text-2xl font-bold">Calendar</h1>
  <div class="calendar-wrapper mt-10 h-full">
    <Calendar {plugins} {options} />
  </div>
</div>
