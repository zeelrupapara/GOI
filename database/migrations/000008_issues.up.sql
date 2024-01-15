-- +migrate Up
CREATE TABLE
    IF NOT EXISTS "issues" (
        "id" VARCHAR (50) PRIMARY KEY NOT NULL,
        "title" TEXT NOT NULL,
        "status" VARCHAR (20) NOT NULL,
        "url" TEXT,
        "number" INTEGER,
        "author_id" VARCHAR (50) NOT NULL,
        "repository_id" VARCHAR (50) NOT NULL,
        "github_closed_at" TIMESTAMP,
        "github_created_at" TIMESTAMP,
        "github_updated_at" TIMESTAMP,
        "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP
    );
