CREATE SCHEMA "newsletter";
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "newsletter"."days" (
	"year" integer NOT NULL,
	"week" integer NOT NULL,
	"day" integer NOT NULL,
	"actual_date" date NOT NULL,
	"lunch" varchar,
	"x_period" varchar,
	"notes" varchar,
	"ap_info" varchar,
	"cc_info" varchar,
	"grade_9" varchar,
	"grade_10" varchar,
	"grade_11" varchar,
	"grade_12" varchar,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT "days_year_week_day_pk" PRIMARY KEY("year","week","day")
);
