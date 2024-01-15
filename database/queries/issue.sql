-- name: InsertIssue :one
INSERT INTO
    "issues" (
        "id",
        "title",
        "status",
        "url",
        "number",
        "author_id",
        "repository_collaborators_id",
        "github_closed_at",
        "github_created_at",
        "github_updated_at"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING issues.id;

-- name: GetIssueByID :one
SELECT issues.id FROM "issues" WHERE issues.id = $1;
