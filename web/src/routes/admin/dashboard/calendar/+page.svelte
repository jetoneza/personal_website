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
  import { applyAction, enhance } from '$app/forms';
  import { invalidateAll } from '$app/navigation';
  import { calendarOptions, getEventColor } from '$lib/utils/calendar';
  import { formatDate, formatInputDate } from '$lib/utils/date';

  // Constants
  import { API_STATUS } from '$lib/constants';

  // Common Components
  import Input from '$lib/components/Input.svelte';

  // Types
  import type { ActionData, PageData } from './$types';

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
    addEvent: (event: CalendarEvent) => void;
    removeEventById: (id: string | number) => void;
  };

  type CalendarEvent = {
    id: string | number;
    title: string;
    start: Date;
    end: Date;
    allDay?: boolean;
  };

  type Event = {
    id: string | number;
    title: string;
    start: string;
    end: string;
    type?: string;
    allDay?: boolean;
    notes?: string;
  };

  type EventResponse = Omit<Event, 'allDay'> & { all_day: boolean };
  type ModalAction = 'new' | 'view';

  // Constants
  const EVENT_TEMP_ID = 'new-event-temp';

  // Props
  export let data: PageData;
  export let form: ActionData;

  // State
  let calendarElement: CalendarElement;
  let openModal = false;
  let modalAction: ModalAction = 'new';
  let activeEvent: Event | null;
  let newEvent: Event = {
    id: EVENT_TEMP_ID,
    title: '(No title)',
    start: '',
    end: '',
    allDay: true,
  };

  const addNewEvent = (start: Date, end: Date, allDay = false) => {
    openModal = true;
    modalAction = 'new';

    newEvent = {
      ...newEvent,
      start: formatInputDate(start),
      end: formatInputDate(end),
      allDay,
    };

    calendarElement.addEvent({
      id: newEvent.id,
      title: newEvent.title,
      start,
      end,
      allDay,
    });
  };

  const handleEventClick = (info: { event: Event }) => {
    activeEvent = data.events.find((event: Event) => event.id === info.event.id);

    openModal = true;
    modalAction = 'view';
  };

  const handleDateClick = (info: DateInfo) => addNewEvent(info.date, info.date, true);
  const handleSelect = (info: DateInfo) => addNewEvent(info.start, info.end);

  const handleCloseModal = () => {
    calendarElement.removeEventById(EVENT_TEMP_ID);
    openModal = false;
  };

  const plugins = [TimeGrid, DayGrid, Interaction];
  $: options = {
    ...calendarOptions,
    eventClick: handleEventClick,
    dateClick: handleDateClick,
    select: handleSelect,
    events: data.events.map(({ id, start, end, all_day, title, type }: EventResponse) => ({
      id,
      start,
      end,
      allDay: all_day,
      title,
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
        modalAction = 'new';
      }}
    >
      Create Event
    </button>
  </div>
  <div class="calendar-wrapper mt-10 h-full">
    <Calendar bind:this={calendarElement} {plugins} {options} />
  </div>
  <Modal
    title={modalAction === 'new' ? 'Create Event' : ''}
    bind:open={openModal}
    on:close={handleCloseModal}
  >
    {#if modalAction === 'new'}
      <!-- TODO: Use floating notification UI -->
      <!-- TODO: Reset form after closing the modal -->
      {#if form?.status === API_STATUS.FAIL}
        <p class="p-4 bg-pink-100 border border-red-500 rounded-lg text-red-500 my-6 text-sm">
          {form?.message}
        </p>
      {/if}

      <form
        method="POST"
        use:enhance={() =>
          async ({ result }) => {
            if (result.type === API_STATUS.SUCCESS) {
              await invalidateAll();
              handleCloseModal();
            }

            applyAction(result);
          }}
      >
        <div class="flex flex-col gap-4">
          <!-- TODO: Support time selection -->
          <Input type="text" name="title" label="Title" placeholder={newEvent.title} required />
          <Input
            type="text"
            name="type"
            label="Type"
            placeholder="Enter type e.g. work, task"
            required
          />
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
            <Input type="date" name="start" value={newEvent.start} label="Start" />
            <Input type="date" name="end" value={newEvent.end} label="End" />
          </div>
          <div class="input-wrapper">
            <label for="all_day" class="block mb-2 text-md font-bold dark:text-white">
              All Day?
            </label>
            <select
              value={newEvent.allDay}
              id="all_day"
              name="all_day"
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
          <div class="textarea-wrapper mt-2">
            <label for="content" class="block mb-2 text-md font-bold dark:text-white">Content</label
            >
            <textarea
              id="notes"
              name="notes"
              rows="10"
              class="w-full rounded-md border border-input bg-background px-3 py-2
    text-sm outline-none focus:outline-zinc-500 font-mono text-black dark:bg-zinc-800
    dark:border-zinc-700 dark:text-white"
            />
          </div>
          <button class="btn mt-2">Create</button>
        </div>
      </form>
    {:else if modalAction === 'view' && activeEvent}
      <div class="flex flex-col gap-2">
        <div class="title flex flex-col gap-1">
          <span class="text-gray-700 font-bold text-2xl dark:text-white">{activeEvent.title}</span>
          <span class="event-date text-sm text-gray-700 dark:text-white">
            {formatDate(activeEvent.start)}
          </span>
        </div>

        <div class="notes text-sm text-gray-600 dark:text-white pt-4">
          {#if activeEvent.notes}
            {activeEvent.notes} 
          {:else}
            <p class="italic">No notes provided.</p>
          {/if}
        </div>
      </div>
    {/if}
  </Modal>
</div>
