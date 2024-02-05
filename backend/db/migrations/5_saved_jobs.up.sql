create table if not exists "saved_jobs" (
  "id" SERIAL,
  "user_id" int NOT NULL REFERENCES "users" ("id") ON DELETE CASCADE,
  "job_id" int NOT NULL REFERENCES "jobs" ("id"),
  "created_at" TIMESTAMPTZ DEFAULT NOW(),
  PRIMARY KEY (user_id, job_id)
);
CREATE INDEX ON saved_jobs (user_id);
