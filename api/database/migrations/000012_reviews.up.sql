-- +migrate Up
CREATE TABLE IF NOT EXISTS "reviews" (
  "id" VARCHAR PRIMARY KEY NOT NULL,
  "reviewer_id" VARCHAR NOT NULL,
  "pr_id" VARCHAR  NOT NULL,
  "status" VARCHAR NOT NULL,
  "github_created_at" TIMESTAMP,
  "github_updated_at" TIMESTAMP,
  "github_submitted_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
