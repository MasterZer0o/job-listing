CREATE table "jobs_content" (
  "id" SERIAL PRIMARY KEY,
  "job_id" int not null REFERENCES "jobs" ("id") ON DELETE CASCADE,
  "content" TEXT,
  "sections" VARCHAR[],
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);

create index on jobs_content (job_id);