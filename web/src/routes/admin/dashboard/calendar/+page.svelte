<script lang="ts">
  // Libraries

  // TODO: Find workaround on sveltekit not reading declared modules in app.d.ts within .svelte files
  // @ts-expect-error: No declaration file for module
  import Calendar from '@event-calendar/core';
  // @ts-expect-error: No declaration file for module
  import TimeGrid from '@event-calendar/time-grid';
  // @ts-expect-error: No declaration file for module
  import DayGrid from '@event-calendar/day-grid';
  // @ts-expect-error: No declaration file for module
  import Interaction from '@event-calendar/interaction';

  // Styles
  import './styles.css';

  // TODO: declare correct specific fields
  type Event = {
    [key: string]: string | number | Date | boolean | undefined;
  };

  // TODO: declare correct specific fields
  type EventTypes = {
    [key: string]: {
      [key: string]: string;
    };
  };

  type DateInfo = {
    date: Date;
    dateStr: string;
  };

  type CalendarElement = {
    addEvent: (event: Event) => void;
  };

  const types: EventTypes = {
    work: {
      backgroundColor: '#0891b2',
    },
    task: {
      backgroundColor: '#14b8a6',
    },
  };

  let calendarElement: CalendarElement;

  const events = [
    {
      id: 2,
      start: new Date('11-21-2023'),
      end: new Date('11-21-2023'),
      allDay: true,
      notes: 'Chill day at work.',
      type: 'work',
      createdAt: new Date(),
    },
    {
      id: 4,
      start: new Date('11-21-2023'),
      end: new Date('11-24-2023'),
      allDay: true,
      notes: 'Project B - Implementation',
      type: 'task',
      createdAt: new Date(),
    },
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

  const handleDateClick = (info: DateInfo) => {
    calendarElement.addEvent({
      id: 'new-temporary',
      start: info.date,
      end: info.date,
      allDay: true,
      title: '(No title)',
    });
  };

  const plugins = [TimeGrid, DayGrid, Interaction];
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
    dateClick: handleDateClick,
    display: 'background',
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
  <div class="actions w-full flex justify-between items-center">
    <h1 class="font-bold text-2xl font-sans-pro">Calendar</h1>
    <a class="btn text-sm" href="/admin/dashboard/blog/create">Create Event</a>
  </div>
  <div class="calendar-wrapper mt-10 h-full">
    <Calendar bind:this={calendarElement} {plugins} {options} />
  </div>
</div>
