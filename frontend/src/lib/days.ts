import dayjs from 'dayjs';
import isoWeek from 'dayjs/plugin/isoWeek';

export declare interface Date {
	year: number;
	week: number;
	day: number;
}

dayjs.extend(isoWeek);

// Get date information from dayjs.
export const getDate = (): Date => {
	const day = dayjs();

	return {
		year: day.year(),
		week: day.isoWeek(),
		day: day.isoWeekday()
	};
};
