-- +migrate Up
CREATE TABLE IF NOT EXISTS "commits" (
  "id" VARCHAR PRIMARY KEY NOT NULL,
  "hash_id" VARCHAR NOT NULL,
  "message" TEXT,
  "branch_id" VARCHAR (50) NOT NULL,
  "pr_id" VARCHAR (50),
  "author_id" VARCHAR (50) NOT NULL,
  "url" TEXT,
  "commit_url" TEXT,
  "github_committed_time" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
