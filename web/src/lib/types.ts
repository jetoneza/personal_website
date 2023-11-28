export type Event = {
  id: string | number;
  title: string;
  start: string;
  end: string;
  type?: string;
  allDay?: boolean;
  notes?: string;
};
