export function formatToCurrency(value: number): string {
  return new Intl.NumberFormat('en-PH', {
    style: 'currency',
    currency: 'PHP',
    minimumIntegerDigits: 1,
    maximumFractionDigits: 2,
  }).format(value);
}
