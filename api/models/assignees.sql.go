// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: assignees.sql

package models

import (
	"context"
	"database/sql"
)

const getAssignedLabalByIssue = `-- name: GetAssignedLabalByIssue :one
SELECT assigned_labals.id
FROM "assigned_labals"
WHERE
    assigned_labals.labal_id = $1
    AND assigned_labals.issue_id = $2
`

type GetAssignedLabalByIssueParams struct {
	LabalID string         `json:"labal_id"`
	IssueID sql.NullString `json:"issue_id"`
}

func (q *Queries) GetAssignedLabalByIssue(ctx context.Context, arg GetAssignedLabalByIssueParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getAssignedLabalByIssue, arg.LabalID, arg.IssueID)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getAssignedLabalByPR = `-- name: GetAssignedLabalByPR :one
SELECT assigned_labals.id
FROM "assigned_labals"
WHERE
    assigned_labals.labal_id = $1
    AND assigned_labals.pr_id = $2
`

type GetAssignedLabalByPRParams struct {
	LabalID string         `json:"labal_id"`
	PrID    sql.NullString `json:"pr_id"`
}

func (q *Queries) GetAssignedLabalByPR(ctx context.Context, arg GetAssignedLabalByPRParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getAssignedLabalByPR, arg.LabalID, arg.PrID)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getAssigneeByIssue = `-- name: GetAssigneeByIssue :one
SELECT assignees.id
FROM "assignees"
WHERE
    assignees.collaborator_id = $1
    AND assignees.issue_id = $2
`

type GetAssigneeByIssueParams struct {
	CollaboratorID string         `json:"collaborator_id"`
	IssueID        sql.NullString `json:"issue_id"`
}

func (q *Queries) GetAssigneeByIssue(ctx context.Context, arg GetAssigneeByIssueParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getAssigneeByIssue, arg.CollaboratorID, arg.IssueID)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getAssigneeByPR = `-- name: GetAssigneeByPR :one
SELECT assignees.id
FROM "assignees"
WHERE
    assignees.collaborator_id = $1
    AND assignees.pr_id = $2
`

type GetAssigneeByPRParams struct {
	CollaboratorID string         `json:"collaborator_id"`
	PrID           sql.NullString `json:"pr_id"`
}

func (q *Queries) GetAssigneeByPR(ctx context.Context, arg GetAssigneeByPRParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getAssigneeByPR, arg.CollaboratorID, arg.PrID)
	var id string
	err := row.Scan(&id)
	return id, err
}

const insertAssignedLabal = `-- name: InsertAssignedLabal :one
INSERT INTO
    "assigned_labals" (
        "id",
        "labal_id",
        "pr_id",
        "issue_id",
        "activity_type"
    )
VALUES ($1, $2, $3, $4, $5) RETURNING assigned_labals.id
`

type InsertAssignedLabalParams struct {
	ID           string         `json:"id"`
	LabalID      string         `json:"labal_id"`
	PrID         sql.NullString `json:"pr_id"`
	IssueID      sql.NullString `json:"issue_id"`
	ActivityType string         `json:"activity_type"`
}

func (q *Queries) InsertAssignedLabal(ctx context.Context, arg InsertAssignedLabalParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertAssignedLabal,
		arg.ID,
		arg.LabalID,
		arg.PrID,
		arg.IssueID,
		arg.ActivityType,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const insertAssignee = `-- name: InsertAssignee :one
INSERT INTO
    "assignees" (
        "id",
        "collaborator_id",
        "pr_id",
        "issue_id",
        "activity_type"
    )
VALUES ($1, $2, $3, $4, $5) RETURNING assignees.id
`

type InsertAssigneeParams struct {
	ID             string         `json:"id"`
	CollaboratorID string         `json:"collaborator_id"`
	PrID           sql.NullString `json:"pr_id"`
	IssueID        sql.NullString `json:"issue_id"`
	ActivityType   string         `json:"activity_type"`
}

func (q *Queries) InsertAssignee(ctx context.Context, arg InsertAssigneeParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertAssignee,
		arg.ID,
		arg.CollaboratorID,
		arg.PrID,
		arg.IssueID,
		arg.ActivityType,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}
