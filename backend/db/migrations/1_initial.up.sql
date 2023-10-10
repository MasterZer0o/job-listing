CREATE TABLE IF NOT EXISTS "users" (
	"id" SERIAL PRIMARY KEY,
	"email" VARCHAR(100) NOT NULL UNIQUE,
	"first_name" VARCHAR(100),
	"last_name" VARCHAR(100),
	"password" VARCHAR(100) NOT NULL,
	"created_at" TIMESTAMPTZ DEFAULT NOW(),
	"updated_at" TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS "sessions" (
	"id" VARCHAR(255) PRIMARY KEY,
	"user_id" INTEGER NOT NULL,
	"created_at" TIMESTAMPTZ DEFAULT NOW(),
	CONSTRAINT "fk_sessions_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

