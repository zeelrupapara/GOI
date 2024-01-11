-- +migrate Up
CREATE TABLE IF NOT EXISTS "branches" (
  "id" VARCHAR (50) PRIMARY KEY NOT NULL,
  "name" VARCHAR (50),
  "url" VARCHAR (50) NOT NULL,
  "repository_id" VARCHAR (50) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
