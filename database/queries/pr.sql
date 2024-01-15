-- name: InsertPR :one
INSERT INTO
    "pull_requests" (
        "id",
        "title",
        "status",
        "url",
        "number",
        "is_draft",
        "branch",
        "author_id",
        "repository_id",
        "github_closed_at",
        "github_merged_at",
        "github_created_at",
        "github_updated_at"
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13
    ) RETURNING pull_requests.id;
-- name: GetPRByID :one
SELECT pull_requests.id
FROM "pull_requests"
WHERE pull_requests.id = $1;
