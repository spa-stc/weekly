// Server-Side only database client.
import postgres from 'postgres';
import { drizzle } from 'drizzle-orm/postgres-js';

const client = postgres(process.env.DB_URL)
const db = drizzle(client)

export { db };