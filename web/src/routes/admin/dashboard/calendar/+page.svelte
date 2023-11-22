<script lang="ts">
  // Libraries
  import { Modal } from 'flowbite-svelte';

  // TODO: Find workaround on sveltekit not reading declared modules in app.d.ts within .svelte files
  // @ts-expect-error: No declaration file for module
  import Calendar from '@event-calendar/core';
  // @ts-expect-error: No declaration file for module
  import TimeGrid from '@event-calendar/time-grid';
  // @ts-expect-error: No declaration file for module
  import DayGrid from '@event-calendar/day-grid';
  // @ts-expect-error: No declaration file for module
  import Interaction from '@event-calendar/interaction';

  // Utils
  import { calendarOptions, getEventColor } from '$lib/utils/calendar';

  // Styles
  import './styles.css';

  // Types
  type DateInfo = {
    date: Date;
    dateStr: string;
    start: Date;
    startStr: string;
    end: Date;
    endStr: string;
  };

  type CalendarElement = {
    addEvent: (event: Event) => void;
  };

  // TODO: declare correct specific fields
  type Event = {
    [key: string]: string | number | Date | boolean | undefined;
  };

  // State
  let openModal = false;
  let calendarElement: CalendarElement;

  // TODO: Remove static data and use real data
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
      notes: 'Christmas Day 🌲🎁🎉',
      type: 'holiday',
      createdAt: new Date(),
    },
  ];

  const addNewEvent = (start: Date, end: Date) => {
    openModal = true;

    calendarElement.addEvent({
      start,
      end,
      id: 'new-temporary',
      allDay: true,
      title: '(No title)',
    });
  };

  const handleEventClick = (info: { event: Event }) => {
    // TODO: Open event modal
    // TODO: Display event
    console.log(info.event.id);
  };

  const handleDateClick = (info: DateInfo) => addNewEvent(info.date, info.date);
  const handleSelect = (info: DateInfo) => addNewEvent(info.start, info.end);

  const plugins = [TimeGrid, DayGrid, Interaction];
  const options = {
    ...calendarOptions,
    eventClick: handleEventClick,
    dateClick: handleDateClick,
    select: handleSelect,
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
    <button
      class="btn text-sm"
      on:click={() => {
        openModal = true;
      }}
    >
      Create Event
    </button>
  </div>
  <div class="calendar-wrapper mt-10 h-full">
    <Calendar bind:this={calendarElement} {plugins} {options} />
  </div>
  <Modal title="Create Event" bind:open={openModal} autoclose>
    <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
      With less than a month to go before the European Union enacts new consumer privacy laws for
      its citizens, companies around the world are updating their terms of service agreements to
      comply.
    </p>
    <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
      The European Union’s General Data Protection Regulation (G.D.P.R.) goes into effect on May 25
      and is meant to ensure a common set of data rights in the European Union. It requires
      organizations to notify users as soon as possible of high-risk data breaches that could
      personally affect them.
    </p>
  </Modal>
</div>
