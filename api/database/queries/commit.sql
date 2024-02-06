-- name: InsertCommit :one
INSERT INTO
    "commits" (
        "id",
        "message",
        "branch_id",
        "author_id",
        "pr_id",
        "url",
        "commit_url",
        "github_committed_time"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING commits.id;

-- name: GetCommitByID :one
SELECT commits.id FROM "commits" WHERE commits.id = $1 and commits.branch_id = $2;
