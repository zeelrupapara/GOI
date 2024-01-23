// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: review.sql

package models

import (
	"context"
	"database/sql"
)

const getReviewByID = `-- name: GetReviewByID :one
SELECT reviews.id FROM "reviews" WHERE reviews.id = $1
`

func (q *Queries) GetReviewByID(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, getReviewByID, id)
	err := row.Scan(&id)
	return id, err
}

const getReviewByPRAndReviewerID = `-- name: GetReviewByPRAndReviewerID :one
SELECT reviews.id
FROM "reviews"
WHERE
    reviews.pr_id = $1
    AND reviews.reviewer_id = $2
`

type GetReviewByPRAndReviewerIDParams struct {
	PrID       string `json:"pr_id"`
	ReviewerID string `json:"reviewer_id"`
}

func (q *Queries) GetReviewByPRAndReviewerID(ctx context.Context, arg GetReviewByPRAndReviewerIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getReviewByPRAndReviewerID, arg.PrID, arg.ReviewerID)
	var id string
	err := row.Scan(&id)
	return id, err
}

const insertReview = `-- name: InsertReview :one
INSERT INTO
    "reviews" (
        "id", "reviewer_id", "pr_id", "status", "github_created_at", "github_updated_at", "github_submitted_at"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING reviews.id
`

type InsertReviewParams struct {
	ID                string       `json:"id"`
	ReviewerID        string       `json:"reviewer_id"`
	PrID              string       `json:"pr_id"`
	Status            string       `json:"status"`
	GithubCreatedAt   sql.NullTime `json:"github_created_at"`
	GithubUpdatedAt   sql.NullTime `json:"github_updated_at"`
	GithubSubmittedAt sql.NullTime `json:"github_submitted_at"`
}

func (q *Queries) InsertReview(ctx context.Context, arg InsertReviewParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertReview,
		arg.ID,
		arg.ReviewerID,
		arg.PrID,
		arg.Status,
		arg.GithubCreatedAt,
		arg.GithubUpdatedAt,
		arg.GithubSubmittedAt,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const updateReview = `-- name: UpdateReview :one
UPDATE "reviews"
SET
    reviewer_id = $2,
    pr_id = $3,
    status = $4,
    github_created_at = $5,
    github_updated_at = $6,
    github_submitted_at = $7
WHERE
    id = $1 RETURNING reviews.id
`

type UpdateReviewParams struct {
	ID                string       `json:"id"`
	ReviewerID        string       `json:"reviewer_id"`
	PrID              string       `json:"pr_id"`
	Status            string       `json:"status"`
	GithubCreatedAt   sql.NullTime `json:"github_created_at"`
	GithubUpdatedAt   sql.NullTime `json:"github_updated_at"`
	GithubSubmittedAt sql.NullTime `json:"github_submitted_at"`
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) (string, error) {
	row := q.db.QueryRowContext(ctx, updateReview,
		arg.ID,
		arg.ReviewerID,
		arg.PrID,
		arg.Status,
		arg.GithubCreatedAt,
		arg.GithubUpdatedAt,
		arg.GithubSubmittedAt,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}
