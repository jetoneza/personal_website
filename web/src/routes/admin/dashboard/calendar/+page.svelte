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
  import { enhance } from '$app/forms';
  import { calendarOptions, getEventColor } from '$lib/utils/calendar';
  import { formatInputDate } from '$lib/utils/date';

  // Common Components
  import Input from '$lib/components/Input.svelte';

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
    removeEventById: (id: string | number) => void;
  };

  // TODO: declare correct specific fields
  type Event = {
    id: string | number;
    title: string;
    start: Date;
    end: Date;
    allDay?: boolean;
  };

  // Constants
  const EVENT_TEMP_ID = 'new-event-temp';

  // State
  let calendarElement: CalendarElement;
  let openModal = false;
  let newEvent: Event = {
    id: EVENT_TEMP_ID,
    title: '(No title)',
    start: new Date(),
    end: new Date(),
  };

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

  const addNewEvent = (start: Date, end: Date, allDay = false) => {
    openModal = true;

    newEvent = {
      ...newEvent,
      start,
      end,
      allDay,
    };

    calendarElement.addEvent(newEvent);
  };

  const handleEventClick = (info: { event: Event }) => {
    // TODO: Open event modal
    // TODO: Display event
    console.log(info.event.id);
  };

  const handleDateClick = (info: DateInfo) => addNewEvent(info.date, info.date, true);
  const handleSelect = (info: DateInfo) => addNewEvent(info.start, info.end);

  const handleCloseModal = () => {
    calendarElement.removeEventById(EVENT_TEMP_ID);
  };

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
  <Modal title="Create Event" bind:open={openModal} on:close={handleCloseModal} autoclose>
    <form method="POST" use:enhance>
      <div class="flex flex-col gap-4">
        <Input type="text" name="title" label="Title" placeholder={newEvent.title} required />
        <Input type="date" name="start" value={formatInputDate(newEvent.start)} label="Start" />
        <Input type="date" name="end" value={formatInputDate(newEvent.end)} label="End" />
        <div class="input-wrapper">
          <label for="published" class="block mb-2 text-md font-bold dark:text-white">
            All Day?
          </label>
          <select
            value={newEvent.allDay}
            id="published"
            name="published"
            class="
              h-10 w-full rounded-md border border-input bg-background px-3
              py-2 text-sm outline-none focus:outline-zinc-500 dark:bg-zinc-800 dark:border-zinc-700
              dark:focus:outline-zinc-600
            "
          >
            <option selected value={false}>No</option>
            <option value={true}>Yes</option>
          </select>
        </div>
        <button class="btn mt-2">Create</button>
      </div>
    </form>
  </Modal>
</div>
