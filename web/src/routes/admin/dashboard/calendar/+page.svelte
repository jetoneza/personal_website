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
  import { formatInputDate } from '$lib/utils/date';

  // Constants
  import { API_STATUS } from '$lib/constants';

  // Common Components
  import EventView from '$lib/components/admin/EventView.svelte';

  // Types
  import type { Event } from '$lib/types';
  import type { ActionData, PageData } from './$types';

  // Styles
  import './styles.css';
  import EventForm from '$lib/components/forms/EventForm.svelte';

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

  const handleDeleteEvent = async (id: string | number) => {
    const response = await fetch(`/events/${id}`, {
      method: 'DELETE',
    });

    const responseData = await response.json();

    if (responseData.status === API_STATUS.SUCCESS) {
      calendarElement.removeEventById(id);
      openModal = false;
    }

    // TODO: Handle error
    // TODO: Add success feedback e.g. alert
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
  <!-- TODO Fix modal dark color -->
  <Modal
    class="dark:bg-zinc-800"
    title={modalAction === 'new' ? 'Create Event' : ''}
    size={modalAction === 'new' ? 'md' : 'sm'}
    bind:open={openModal}
    on:close={handleCloseModal}
  >
    {#if modalAction === 'new'}
      <!-- TODO: Use floating notification UI -->
      {#if form?.status === API_STATUS.FAIL}
        <p class="p-4 bg-pink-100 border border-red-500 rounded-lg text-red-500 my-6 text-sm">
          {form?.message}
        </p>
      {/if}

      <form
        method="POST"
        class="text-slate-800 dark:text-white"
        use:enhance={() =>
          async ({ result }) => {
            if (result.type === API_STATUS.SUCCESS) {
              await invalidateAll();
              handleCloseModal();
            }

            applyAction(result);
          }}
      >
        <EventForm start={newEvent.start} end={newEvent.end} allDay={newEvent.allDay} />
      </form>
    {:else if modalAction === 'view' && activeEvent}
      <EventView {activeEvent} onDeleteClick={handleDeleteEvent} />
    {/if}
  </Modal>
</div>
