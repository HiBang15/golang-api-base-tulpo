CREATE TABLE IF NOT EXISTS  "permissions" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz,
  "updated_at" timestamptz DEFAULT (now())
);