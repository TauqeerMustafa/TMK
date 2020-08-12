export interface MetricCardProps {
  title: string;
  value: string | number;
  unit?: string;
  changePercent?: number;
  status?: 'good' | 'warn' | 'bad';
}

export function formatMetric(val: number, decimals: number = 2): string {
  return val.toLocaleString(undefined, {
    minimumFractionDigits: decimals,
    maximumFractionDigits: decimals
  });
}

export const MetricCard = (props: MetricCardProps) => {
  const { title, value, unit, changePercent, status = 'good' } = props;
  return {
    title,
    formattedValue: `${value}${unit ? ' ' + unit : ''}`,
    status,
    trend: changePercent ? `${changePercent > 0 ? '+' : ''}${changePercent}%` : 'stable'
  };
};
