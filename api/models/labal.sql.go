// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: labal.sql

package models

import (
	"context"
	"database/sql"
)

const getLabalByID = `-- name: GetLabalByID :one
SELECT labals.id FROM "labals" WHERE labals.id = $1
`

func (q *Queries) GetLabalByID(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, getLabalByID, id)
	err := row.Scan(&id)
	return id, err
}

const insertLabal = `-- name: InsertLabal :one
INSERT INTO
    "labals" ("id", "name")
VALUES ($1, $2) RETURNING labals.id
`

type InsertLabalParams struct {
	ID   string         `json:"id"`
	Name sql.NullString `json:"name"`
}

func (q *Queries) InsertLabal(ctx context.Context, arg InsertLabalParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertLabal, arg.ID, arg.Name)
	var id string
	err := row.Scan(&id)
	return id, err
}
