-- +migrate Up
CREATE TABLE IF NOT EXISTS "organizations" (
  "id" VARCHAR (50) PRIMARY KEY NOT NULL,
  "login" VARCHAR (100) NOT NULL,
  "name" TEXT NOT NULL DEFAULT NULL,
  "email" VARCHAR,
  "location" TEXT,
  "description" TEXT,
  "url" TEXT,
  "avatar_url" TEXT,
  "website_url" TEXT,
  "github_updated_at" TIMESTAMP,
  "github_created_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

