import type { days } from './schema';

type Day = typeof days.$inferSelect;

interface Announcement {
	title: string | null;
	author: string | null;
	content: string | null;
	date: string;
}

export type { Day, Announcement };
