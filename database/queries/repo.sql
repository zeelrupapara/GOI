-- name: InsertRepo :one
INSERT INTO "repositories" (
        "id",
        "name",
        "is_private",
        "default_branch",
        "url",
        "homepage_url",
        "open_issues",
        "closed_issues",
        "open_prs",
        "closed_prs",
        "merged_prs",
        "organization_collaborator_id",
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
        $13,
        $14
    )
RETURNING repositories.id;
-- name: GetRepoByID :one
SELECT repositories.id
FROM "repositories"
WHERE repositories.id = $1;
