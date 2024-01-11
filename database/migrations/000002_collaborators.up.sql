-- +migrate Up
CREATE TABLE IF NOT EXISTS "collaborators" (
  "id" VARCHAR (50) PRIMARY KEY NOT NULL,
  "name" VARCHAR (100),
  "username" varchar (100) NOT NULL,
  "email" VARCHAR,
  "url" TEXT,
  "avatar_url" TEXT,
  "website_url" VARCHAR,
  "github_created_at" TIMESTAMP,
  "github_updated_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);


