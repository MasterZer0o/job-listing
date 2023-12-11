create table if not exists "job_experience_levels" (
  "id" SERIAL PRIMARY KEY UNIQUE,
  "job_id" INTEGER NOT NULL REFERENCES "jobs" ("id") ON DELETE CASCADE,
  "exp_level_id" INTEGER NOT NULL REFERENCES "experience_levels" ("id") ON DELETE CASCADE,
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX ON job_experience_levels ("job_id");


create table if not exists "job_employment_types" (
  "id" SERIAL PRIMARY KEY UNIQUE,
  "job_id" INTEGER NOT NULL REFERENCES "jobs" ("id") ON DELETE CASCADE,
  "emp_type_id" INTEGER NOT NULL REFERENCES "employment_types" ("id") ON DELETE CASCADE,
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX ON job_employment_types ("job_id");

create table if not exists "job_tech_skills" (
  "id" SERIAL PRIMARY KEY UNIQUE,
  "job_id" INTEGER NOT NULL REFERENCES "jobs" ("id") ON DELETE CASCADE,
  "tech_skill_id" INTEGER NOT NULL REFERENCES "tech_skills" ("id") ON DELETE CASCADE,
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX ON job_tech_skills ("job_id");

create table if not exists "job_types_of_work" (
  "id" SERIAL PRIMARY KEY UNIQUE,
  "job_id" INTEGER NOT NULL REFERENCES "jobs" ("id") ON DELETE CASCADE,
  "type_of_work_id" INTEGER NOT NULL REFERENCES "types_of_work" ("id") ON DELETE CASCADE,
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX ON job_types_of_work (job_id);

create table if not exists "job_work_modes" (
  "id" SERIAL PRIMARY KEY UNIQUE,
  "job_id" INTEGER NOT NULL REFERENCES "jobs" ("id") ON DELETE CASCADE,
  "work_mode_id" INTEGER NOT NULL REFERENCES "work_modes" ("id") ON DELETE CASCADE,
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX ON job_work_modes (job_id);



