CREATE TABLE IF NOT EXISTS  "role_permission" (
  "role_id" int NOT NULL,
  "permission_id" int NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz,
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id") DEFERRABLE INITIALLY DEFERRED;