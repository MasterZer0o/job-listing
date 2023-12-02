create table if not exists "experience_levels" (
  "id" smallint PRIMARY KEY,
  "name" VARCHAR(50),
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);

create table if not exists "employment_types" (
  "id" smallint PRIMARY KEY,
  "name" VARCHAR(50),
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);

create table if not exists "tech_skills" (
  "id" smallint PRIMARY KEY,
  "name" VARCHAR(50),
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);

create table if not exists "types_of_work" (
  "id" smallint PRIMARY KEY,
  "name" VARCHAR(50),
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);

create table if not exists "work_modes" (
  "id" smallint PRIMARY KEY,
  "name" VARCHAR(50),
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);



