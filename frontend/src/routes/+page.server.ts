import { getDate } from '$lib/days';
import { days } from '$lib/schema';
import { db } from '$lib/server/database';
import { eq, and, gt, lt } from 'drizzle-orm';
import type { PageServerLoad } from './$types';
import { error } from '@sveltejs/kit';

// Load data for the homepage.
export const load: PageServerLoad = async () => {
	const appweek = getDate();
	const week = await db
		.select()
		.from(days)
		.where(
			and(
				eq(days.year, appweek.year),
				eq(days.week, appweek.week),
				gt(days.weekday, 1),
				lt(days.weekday, 7)
			)
		);

	if (!(week.length > 0)) {
		error(404, 'week not found');
	}

	return {
		week: week
	};
};
