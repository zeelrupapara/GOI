-- +migrate Up
CREATE TABLE IF NOT EXISTS "reviews" (
  "id" VARCHAR PRIMARY KEY NOT NULL,
  "reviewer_id" VARCHAR NOT NULL,
  "pr_id" VARCHAR  NOT NULL,
  "status" VARCHAR NOT NULL,
  "github_created_at" VARCHAR,
  "github_updated_at" VARCHAR,
  "github_submitted_at" VARCHAR,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
