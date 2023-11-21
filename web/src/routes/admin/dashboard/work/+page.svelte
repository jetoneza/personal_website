<script lang="ts">
  // Libraries
  import Calendar from '@event-calendar/core';
  import TimeGrid from '@event-calendar/time-grid';
  import DayGrid from '@event-calendar/day-grid';

  // Styles
  import './styles.css'

  const events = [
    {
      id: 1,
      start: new Date('11-20-2023'),
      end: new Date('11-20-2023'),
      allDay: true,
      notes: 'First day at Bookipi.',
      type: 'work',
      createdAt: new Date(),
    },
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

    const colors: {
      [key: string]: any;
    } = {
      work: {
        backgroundColor: '#0891b2',
      },
    };

    return colors[type];
  };

  const handleEventClick = (info: any) => {
    console.log(info);
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
  <h1 class="text-2xl font-bold">Work Logs</h1>
  <div class="calendar-wrapper mt-10 h-full">
    <Calendar {plugins} {options} />
  </div>
</div>
