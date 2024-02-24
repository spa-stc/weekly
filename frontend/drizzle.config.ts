import { defineConfig } from 'drizzle-kit';

export default defineConfig({
	out: '../internal/res/migrations',
	schema: './src/lib/schema.ts',
	driver: 'pg',
	dbCredentials: {
		connectionString: process.env.DB_URL
	},
	verbose: true,
	strict: true,
	schemaFilter: ['public', 'newsletter']
});
