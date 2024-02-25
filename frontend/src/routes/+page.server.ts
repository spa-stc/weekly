import { getDate } from '$lib/days';
import { days } from '$lib/schema';
import { db } from '$lib/server/database';
import { eq, and, gte, lte } from 'drizzle-orm';
import type { PageServerLoad } from './$types';

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
				gte(days.weekday, 1),
				lte(days.weekday, 5)
			)
		);

	return {
		week: week
	};
};
