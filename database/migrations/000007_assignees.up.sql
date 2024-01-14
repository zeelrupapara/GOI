-- +migrate Up
CREATE TABLE
    IF NOT EXISTS "assignees" (
        "id" VARCHAR (50) PRIMARY KEY NOT NULL,
        "collaborator_id" VARCHAR (50) NOT NULL,
        "activity_id" VARCHAR (50) NOT NULL,
        "activity_type" VARCHAR (20) NOT NULL,
        "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP
    );
