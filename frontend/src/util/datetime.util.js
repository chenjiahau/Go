import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

export const formatDateTime = (date) => {
  dayjs.extend(utc);
  return dayjs.utc(date).format('YYYY-MM-DD HH:mm:ss');
}
