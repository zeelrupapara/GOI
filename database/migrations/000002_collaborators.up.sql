-- +migrate Up
CREATE TABLE IF NOT EXISTS "collaborators" (
  "id" VARCHAR (50) PRIMARY KEY NOT NULL,
  "name" VARCHAR (50),
  "username" varchar (50) NOT NULL,
  "email" VARCHAR (50),
  "avatar_url" VARCHAR (50),
  "github_url" VARCHAR (50) NOT NULL,
  "website_url" VARCHAR,
  "github_created_at" timestamp,
  "github_updated_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp
);


