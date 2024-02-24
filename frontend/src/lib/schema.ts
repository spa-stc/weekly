import { integer, timestamp, date, pgSchema, varchar, primaryKey } from 'drizzle-orm/pg-core';

// Required app schema.
export const appSchema = pgSchema('newsletter');

// Days table.
export const days = appSchema.table(
	'days',
	{
		// Fetch Related Information.
		year: integer('year').notNull(),
		week: integer('week').notNull(),
		weekday: integer('day').notNull(),
		date: date('actual_date').notNull(),

		// Application Layer Information.
		lunch: varchar('lunch'),
		x_period: varchar('x_period'),
		notes: varchar('notes'),
		ap_info: varchar('ap_info'),
		cc_info: varchar('cc_info'),
		grade9: varchar('grade_9'),
		grade10: varchar('grade_10'),
		grade11: varchar('grade_11'),
		grade12: varchar('grade_12'),

		// Other Stored Information.
		created_at: timestamp('created_at').notNull().defaultNow(),
		updated_at: timestamp('updated_at').notNull().defaultNow()
	},
	(days) => {
		return {
			datePrimaryKey: primaryKey({
				columns: [days.year, days.week, days.weekday]
			})
		};
	}
);
