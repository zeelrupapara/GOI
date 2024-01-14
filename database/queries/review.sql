-- name: InsertReview :one
INSERT INTO
    "reviews" (
        "id",
        "reviewer_id",
        "pr_id",
        "status",
        "github_created_at",
        "github_updated_at",
        "github_submitted_at"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING reviews.id;

-- name: GetReviewByID :one
SELECT reviews.id FROM "reviews" WHERE reviews.id = $1;
