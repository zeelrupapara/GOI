-- name: InsertReview :one
INSERT INTO
    "reviews" (
        "id", "reviewer_id", "pr_id", "status", "github_created_at", "github_updated_at", "github_submitted_at"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING reviews.id;

-- name: GetReviewByID :one
SELECT reviews.id FROM "reviews" WHERE reviews.id = $1;

-- name: GetReviewByPRAndReviewerID :one
SELECT reviews.id
FROM "reviews"
WHERE
    reviews.pr_id = $1
    AND reviews.reviewer_id = $2;

-- name: UpdateReview :one
UPDATE "reviews"
SET
    id = $2,
    reviewer_id = $3,
    pr_id = $4,
    status = $5,
    github_created_at = $6,
    github_updated_at = $7,
    github_submitted_at = $8
WHERE
    id = $1 RETURNING reviews.id;
