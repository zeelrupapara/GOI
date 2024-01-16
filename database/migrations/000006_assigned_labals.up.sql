-- +migrate Up
CREATE TABLE IF NOT EXISTS "assigned_labals" (
  "id" VARCHAR (50) PRIMARY KEY NOT NULL,
  "labal_id" VARCHAR (50) NOT NULL,
  "pr_id" VARCHAR (50),
  "issue_id" VARCHAR (50),
  "activity_type" VARCHAR (20) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
