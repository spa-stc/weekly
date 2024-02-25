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
	let day = dayjs().utcOffset(-6);

	// Allow pre-fetching of the next week starting on Sunday.
	if (day.isoWeekday() == 7) {
		day = day.add(1, 'day');
	}

	return {
		year: day.year(),
		week: day.isoWeek(),
		day: day.isoWeekday()
	};
};
