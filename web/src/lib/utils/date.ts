const monthNames = [
  'January',
  'February',
  'March',
  'April',
  'May',
  'June',
  'July',
  'August',
  'September',
  'October',
  'November',
  'December',
];

export const formatDate = (rawDate: string): string => {
  const date = new Date(rawDate);
  return `${monthNames[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`;
};

export const formatInputDate = (date: Date): string => {
  return [date.getFullYear(), padTo2Digits(date.getMonth() + 1), padTo2Digits(date.getDate())].join(
    '-'
  );
};

const padTo2Digits = (num: number): string => {
  return num.toString().padStart(2, '0');
};
