<script lang="ts">
  import { PUBLIC_MONTHLY_RATE } from '$env/static/public';

  // Components
  import {
    Table,
    TableHead,
    TableHeadCell,
    TableBody,
    TableBodyRow,
    TableBodyCell,
  } from 'flowbite-svelte';

  // Types
  import type { Event } from '$lib/types';

  // Utils
  import { formatDateObject, getWorkDaysInCurrentMonth, isSameDay } from '$lib/utils/date';
  import { formatToCurrency } from '$lib/utils/currency';

  const HOLIDAY_RATE = 1.3;
  const MONTHLY_RATE = parseFloat(PUBLIC_MONTHLY_RATE as string) || 0;

  export let events: Event[] = [];

  $: daysWorked = events.reduce(
    (acc: number, curr: Event) => (curr.type === 'work' ? acc + 1 : acc),
    0
  );

  $: regularWork = daysWorked;
  $: holidayWork = 0;

  $: specialHolidays = events
    .map((event: Event) => {
      if (event.type === 'special-holiday') {
        return event;
      }
    })
    .filter(Boolean);

  $: {
    holidayWork = 0;

    specialHolidays.forEach((event: Event | undefined) => {
      const worked = events.find(
        (e: Event) => event && e.type == 'work' && isSameDay(e.start, event.start)
      );

      if (worked) {
        regularWork -= 1;
        holidayWork += 1;
      }
    });
  }

  let workDays = getWorkDaysInCurrentMonth();
  let dailyRate = MONTHLY_RATE / workDays;

  $: regularWorkDays = workDays - specialHolidays.length;
  $: regularTotal = dailyRate * regularWork;
  $: holidayTotal = dailyRate * holidayWork * HOLIDAY_RATE;
  $: total = regularTotal + holidayTotal;
</script>

<div
  class="mt-10 md:space-y-5 p-6 bg-white border border-zinc-200 rounded-lg shadow dark:bg-zinc-900 dark:border-zinc-800"
>
  <div class="font-semibold">
    Days worked this month,
    <span class="font-bold">
      {formatDateObject(new Date(), 'MM, YYYY')}
    </span>
  </div>
  <div class="work-days">
    <div class="text-6xl font-bold font-sans-pro">{daysWorked}/{workDays}</div>
  </div>

  <Table class="hidden md:block">
    <TableHead class="dark:bg-neutral-700 dark:border-neutral-700">
      <TableHeadCell class="dark:text-gray-100">Work days</TableHeadCell>
      <TableHeadCell class="dark:text-gray-100">Regular</TableHeadCell>
      <TableHeadCell class="dark:text-gray-100">Special</TableHeadCell>
      <TableHeadCell class="dark:text-gray-100">Days worked</TableHeadCell>
    </TableHead>
    <TableBody>
      <TableBodyRow class="dark:bg-neutral-800 dark:border-neutral-700">
        <TableBodyCell>{workDays}</TableBodyCell>
        <TableBodyCell>{regularWorkDays}</TableBodyCell>
        <TableBodyCell>{specialHolidays.length}</TableBodyCell>
        <TableBodyCell>{daysWorked}/{workDays}</TableBodyCell>
      </TableBodyRow>
    </TableBody>
  </Table>

  <div class="hidden md:block font-semibold">
    Breakdown of the monthly rate of <span class="font-bold">{formatToCurrency(MONTHLY_RATE)}</span
    >:
  </div>

  <Table class="hidden md:block text-right">
    <TableHead class="dark:bg-neutral-700 dark:border-neutral-700">
      <TableHeadCell class="text-left dark:text-gray-100">Type</TableHeadCell>
      <TableHeadCell class="text-left dark:text-gray-100">Description</TableHeadCell>
      <TableHeadCell class="text-left dark:text-gray-100">Notes</TableHeadCell>
      <TableHeadCell>Rate</TableHeadCell>
      <TableHeadCell># days worked</TableHeadCell>
      <TableHeadCell>Total</TableHeadCell>
    </TableHead>
    <TableBody>
      <TableBodyRow class="dark:bg-neutral-800 dark:border-neutral-700">
        <TableBodyCell class="text-left">Regular</TableBodyCell>
        <TableBodyCell class="text-left">
          Daily rate for {workDays} workdays this month. <br />
        </TableBodyCell>
        <TableBodyCell class="text-left">
          Worked {regularWork} out of {regularWorkDays} regular workday{regularWorkDays > 1
            ? 's'
            : ''}.
        </TableBodyCell>
        <TableBodyCell>{formatToCurrency(dailyRate)}</TableBodyCell>
        <TableBodyCell>{regularWork}</TableBodyCell>
        <TableBodyCell>{formatToCurrency(regularTotal)}</TableBodyCell>
      </TableBodyRow>
      <TableBodyRow class="dark:bg-neutral-800 dark:border-neutral-700">
        <TableBodyCell class="text-left">Special holiday</TableBodyCell>
        <TableBodyCell class="text-left">
          130% rate for special holidays. <br />
        </TableBodyCell>
        <TableBodyCell class="text-left">
          Worked {holidayWork} out of {specialHolidays.length} holiday{specialHolidays.length > 1
            ? 's'
            : ''}.
        </TableBodyCell>
        <TableBodyCell>{formatToCurrency(dailyRate * HOLIDAY_RATE)}</TableBodyCell>
        <TableBodyCell>{holidayWork}</TableBodyCell>
        <TableBodyCell>{formatToCurrency(holidayTotal)}</TableBodyCell>
      </TableBodyRow>
      <TableBodyRow class="dark:bg-neutral-800 dark:border-neutral-700">
        <TableBodyCell />
        <TableBodyCell />
        <TableBodyCell />
        <TableBodyCell />
        <TableBodyCell class="font-bold">TOTAL:</TableBodyCell>
        <TableBodyCell class="font-bold">{formatToCurrency(total)}</TableBodyCell>
      </TableBodyRow>
    </TableBody>
  </Table>
</div>
