-- +migrate Up
CREATE TABLE IF NOT EXISTS "organizations" (
  "id" CHAR (20) PRIMARY KEY,
  "name" VARCHAR (100) NOT NULL,
  "full_name" TEXT,
  "description" TEXT,
  "email" VARCHAR,
  "url" TEXT,
  "avatar_url" TEXT,
  "website_url" TEXT,
  "github_updated_at" TIMESTAMP,
  "github_created_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
