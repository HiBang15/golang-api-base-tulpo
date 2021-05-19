CREATE TABLE IF NOT EXISTS  "activities" (
  "id" SERIAL PRIMARY KEY,
  "url" varchar NOT NULL,
  "method" varchar,
  "url_regex" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz,
  "updated_at" timestamptz DEFAULT (now())
);