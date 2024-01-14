-- +migrate Up
CREATE TABLE IF NOT EXISTS "repositories" (
  "id" VARCHAR (50) PRIMARY KEY NOT NULL,
  "name" VARCHAR (100),
  "is_private" BOOLEAN,
  "default_branch" VARCHAR,
  "url" TEXT,
  "homepage_url" TEXT,
  "open_issues" INTEGER,
  "closed_issues" INTEGER,
  "open_prs" INTEGER,
  "closed_prs" INTEGER,
  "merged_prs" INTEGER,
  "organization_collaborator_id" VARCHAR (50) NOT NULL,
  "github_created_at" TIMESTAMP,
  "github_updated_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
