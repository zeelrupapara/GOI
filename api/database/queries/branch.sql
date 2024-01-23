-- name: InsertBranch :one
INSERT INTO
    "branches" (
        "id",
        "name",
        "url",
        "repository_id"
    )
VALUES ($1, $2, $3, $4) RETURNING branches.id;

-- name: GetBranchByID :one
SELECT branches.id
FROM "branches"
WHERE
    branches.repository_id = $1
    AND branches.name = $2;
