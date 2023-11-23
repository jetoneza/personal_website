import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
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

  try {
    return { events };
  } catch (_) {
    return { events: [] };
  }
};
