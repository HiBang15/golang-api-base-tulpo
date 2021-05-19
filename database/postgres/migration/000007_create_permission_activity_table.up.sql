CREATE TABLE IF NOT EXISTS  "permission_activity" (
  "permission_id" int NOT NULL,
  "activity_id" int NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz,
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE "permission_activity" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "permission_activity" ADD FOREIGN KEY ("activity_id") REFERENCES "activities" ("id") DEFERRABLE INITIALLY DEFERRED;