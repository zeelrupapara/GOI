-- name: GetOrganizationList :many
SELECT *
FROM "organizations";
-- name: InsertOrganization :one
INSERT INTO "organizations" (
        "id",
        "login",
        "name",
        "email",
        "location",
        "description",
        "url",
        "avatar_url",
        "website_url",
        "github_created_at",
        "github_updated_at"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING organizations.id;
-- name: GetOrganizationByLogin :one
SELECT organizations.id
FROM "organizations"
WHERE login = $1;
