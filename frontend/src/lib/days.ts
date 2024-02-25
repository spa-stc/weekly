import dayjs from 'dayjs';
import isoWeek from 'dayjs/plugin/isoWeek';
import utc from 'dayjs/plugin/utc';

export declare interface Date {
	year: number;
	week: number;
	day: number;
}

dayjs.extend(isoWeek);
dayjs.extend(utc);

// Get date information from dayjs.
export const getDate = (): Date => {
	const day = dayjs().utcOffset(-6);

	return {
		year: day.year(),
		week: day.isoWeek(),
		day: day.isoWeekday()
	};
};
