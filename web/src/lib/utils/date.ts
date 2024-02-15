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
      date.getDate()
    )}`;
  }

  if (format === 'MM, YYYY') {
    return `${monthNames[date.getMonth()]}, ${date.getFullYear()}`;
  }

  return date.toLocaleDateString();
}

export function getWorkDaysInCurrentMonth(): number {
  const currentDate = new Date();
  const currentMonth = currentDate.getMonth();
  const currentYear = currentDate.getFullYear();

  const firstDay = new Date(currentYear, currentMonth, 1);
  const lastDay = new Date(currentYear, currentMonth + 1, 0);

  let workDays = 0;

  for (let day = firstDay; day <= lastDay; day.setDate(day.getDate() + 1)) {
    if (day.getDay() !== 0 && day.getDay() !== 6) {
      workDays++;
    }
  }

  return workDays;
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

export function isInCurrentMonth(dateString: string) {
  const inputDate = new Date(dateString);
  const now = new Date();
  return inputDate.getMonth() === now.getMonth() && inputDate.getFullYear() === now.getFullYear();
}

function padTo2Digits(num: number): string {
  return num.toString().padStart(2, '0');
}
