// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: collaborator.sql

package models

import (
	"context"
	"database/sql"
)

const getMemberByLogin = `-- name: GetMemberByLogin :one
SELECT collaborators.id
FROM "collaborators"
WHERE collaborators.login = $1
`

func (q *Queries) GetMemberByLogin(ctx context.Context, login string) (string, error) {
	row := q.db.QueryRowContext(ctx, getMemberByLogin, login)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getMemberDetailsByLogin = `-- name: GetMemberDetailsByLogin :one
SELECT id, login, name, email, url, avatar_url, website_url, github_created_at, github_updated_at, created_at, updated_at, deleted_at 
FROM "collaborators" 
WHERE collaborators.login = $1
`

func (q *Queries) GetMemberDetailsByLogin(ctx context.Context, login string) (Collaborator, error) {
	row := q.db.QueryRowContext(ctx, getMemberDetailsByLogin, login)
	var i Collaborator
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Name,
		&i.Email,
		&i.Url,
		&i.AvatarUrl,
		&i.WebsiteUrl,
		&i.GithubCreatedAt,
		&i.GithubUpdatedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getMemberIDs = `-- name: GetMemberIDs :many
SELECT DISTINCT
    collaborators.id
FROM "collaborators"
`

func (q *Queries) GetMemberIDs(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getMemberIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMembers = `-- name: GetMembers :many
SELECT id, login, name, email, url, avatar_url, website_url, github_created_at, github_updated_at, created_at, updated_at, deleted_at 
FROM "collaborators" 
ORDER BY collaborators.login
`

func (q *Queries) GetMembers(ctx context.Context) ([]Collaborator, error) {
	rows, err := q.db.QueryContext(ctx, getMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collaborator
	for rows.Next() {
		var i Collaborator
		if err := rows.Scan(
			&i.ID,
			&i.Login,
			&i.Name,
			&i.Email,
			&i.Url,
			&i.AvatarUrl,
			&i.WebsiteUrl,
			&i.GithubCreatedAt,
			&i.GithubUpdatedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertMember = `-- name: InsertMember :one
INSERT INTO "collaborators" (
        "id",
        "login",
        "name",
        "email",
        "url",
        "avatar_url",
        "website_url",
        "github_created_at",
        "github_updated_at"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING collaborators.id
`

type InsertMemberParams struct {
	ID              string         `json:"id"`
	Login           string         `json:"login"`
	Name            sql.NullString `json:"name"`
	Email           sql.NullString `json:"email"`
	Url             sql.NullString `json:"url"`
	AvatarUrl       sql.NullString `json:"avatar_url"`
	WebsiteUrl      sql.NullString `json:"website_url"`
	GithubCreatedAt sql.NullTime   `json:"github_created_at"`
	GithubUpdatedAt sql.NullTime   `json:"github_updated_at"`
}

func (q *Queries) InsertMember(ctx context.Context, arg InsertMemberParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertMember,
		arg.ID,
		arg.Login,
		arg.Name,
		arg.Email,
		arg.Url,
		arg.AvatarUrl,
		arg.WebsiteUrl,
		arg.GithubCreatedAt,
		arg.GithubUpdatedAt,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}
