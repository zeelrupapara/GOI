-- name: InsertMember :one
INSERT INTO "collaborators" (
        "id",
        "login",
        "name",
        "email",
        "url",
        "avatar_url",
        "website_url",
        "github_created_at",
        "github_updated_at"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING collaborators.id;
-- name: GetMemberByLogin :one
SELECT collaborators.id
FROM "collaborators"
WHERE collaborators.login = $1;


-- name: GetMembers :many
SELECT * 
FROM "collaborators" 
ORDER BY collaborators.login;

-- name: GetMemberIDs :many
SELECT DISTINCT
    collaborators.id
FROM "collaborators";
