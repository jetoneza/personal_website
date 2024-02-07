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

export function formatDate(rawDate: string): string {
  const date = new Date(rawDate);
  return `${monthNames[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`;
}

export function formatInputDate(date: Date): string {
  return [date.getFullYear(), padTo2Digits(date.getMonth() + 1), padTo2Digits(date.getDate())].join(
    '-'
  );
}

export function formatDateObject(date: Date, format?: string) {
  if (format === 'YYYYMMDD') {
    return `${date.getFullYear()}${padTo2Digits(date.getMonth() + 1)}${padTo2Digits(
      date.getDay() + 1
    )}`;
  }

  return date.toLocaleDateString();
}

export function isSameDay(dateString1: string, dateString2: string): boolean {
  const date1 = new Date(dateString1);
  const date2 = new Date(dateString2);

  return (
    date1.getFullYear() === date2.getFullYear() &&
    date1.getMonth() === date2.getMonth() &&
    date1.getDate() === date2.getDate()
  );
}

function padTo2Digits(num: number): string {
  return num.toString().padStart(2, '0');
}
