-- +migrate Up
CREATE TABLE IF NOT EXISTS "organization_collaborators" (
  "id" VARCHAR (50) PRIMARY KEY NOT NULL,
  "organization_id" VARCHAR (50) NOT NULL,
  "collaborator_id" VARCHAR (50) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
