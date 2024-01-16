-- +migrate Up
CREATE TABLE
    IF NOT EXISTS "repository_collaborators" (
        "id" VARCHAR (50) PRIMARY KEY NOT NULL,
        "repo_id" VARCHAR (50) NOT NULL,
        "organization_collaborator_id" VARCHAR (50) NOT NULL
    );
