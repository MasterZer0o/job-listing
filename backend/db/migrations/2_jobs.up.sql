CREATE TABLE IF NOT EXISTS "companies" (
	"id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255),
  "image" TEXT,
	"created_at" TIMESTAMPTZ DEFAULT NOW(),
	"updated_at" TIMESTAMPTZ
);
CREATE UNIQUE INDEX id_idx ON companies (id);


CREATE TABLE IF NOT EXISTS "currencies" (
	"id" SERIAL PRIMARY KEY,
  "short_name" VARCHAR(10),
	"created_at" TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS "jobs" (
	"id" SERIAL PRIMARY KEY,
	"created_at" TIMESTAMPTZ DEFAULT NOW(),
  "location" VARCHAR(255),
  "remote_available" BOOLEAN,
  "company_id" INTEGER,
  "level" VARCHAR(20),
  "title" VARCHAR(255),
  "salary_from" INTEGER,
  "salary_to" INTEGER,
  "currency_id" INTEGER,
	"updated_at" TIMESTAMPTZ,
	CONSTRAINT "fk_company_id" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE,
	CONSTRAINT "fk_currency_id" FOREIGN KEY ("currency_id") REFERENCES "currencies" ("id")
);
CREATE UNIQUE INDEX id_idx ON jobs (id);


CREATE TABLE IF NOT EXISTS "known_skills" (
  "name" VARCHAR(100)
);


CREATE TABLE IF NOT EXISTS "job_skills" (
	"id" SERIAL PRIMARY KEY,
  "name" VARCHAR(50),
  "job_id" INTEGER,
	"created_at" TIMESTAMPTZ DEFAULT NOW(),
	"updated_at" TIMESTAMPTZ,
	CONSTRAINT "fk_job_id" FOREIGN KEY ("job_id") REFERENCES "jobs" ("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX job_id_idx ON companies (id);
