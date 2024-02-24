// Server-Side only database client.
import postgres from 'postgres';
import { drizzle } from 'drizzle-orm/postgres-js';
import { DB_URL } from '$env/static/private';

const client = postgres(DB_URL);
const db = drizzle(client);

export { db };
