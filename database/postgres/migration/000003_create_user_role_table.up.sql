CREATE TABLE IF NOT EXISTS  "user_role" (
  "user_id" bigint NOT NULL,
  "role_id" int NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz,
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE "user_role" ADD FOREIGN KEY ("user_id") REFERENCES "user_account" ("id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "user_role" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id") DEFERRABLE INITIALLY DEFERRED;