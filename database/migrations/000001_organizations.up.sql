-- +migrate Up
CREATE TABLE IF NOT EXISTS "organizations" (
  "id" CHAR (20) PRIMARY KEY,
  "name" VARCHAR (50) NOT NULL,
  "full_name" VARCHAR (100),
  "description" TEXT,
  "email" VARCHAR,
  "avatar_url" VARCHAR,
  "website_url" VARCHAR,
  "github_url" VARCHAR,
  "github_updated_at" TIMESTAMP,
  "github_created_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
