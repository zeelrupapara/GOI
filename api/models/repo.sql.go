// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: repo.sql

package models

import (
	"context"
	"database/sql"
)

const getRepoByID = `-- name: GetRepoByID :one
SELECT repositories.id
FROM "repositories"
WHERE repositories.id = $1
`

func (q *Queries) GetRepoByID(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, getRepoByID, id)
	err := row.Scan(&id)
	return id, err
}

const getRepoDetailsByID = `-- name: GetRepoDetailsByID :one
select id, name, is_private, default_branch, url, homepage_url, open_issues, closed_issues, open_prs, closed_prs, merged_prs, github_created_at, github_updated_at, created_at, updated_at, deleted_at from repositories where id = $1
`

func (q *Queries) GetRepoDetailsByID(ctx context.Context, id string) (Repository, error) {
	row := q.db.QueryRowContext(ctx, getRepoDetailsByID, id)
	var i Repository
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsPrivate,
		&i.DefaultBranch,
		&i.Url,
		&i.HomepageUrl,
		&i.OpenIssues,
		&i.ClosedIssues,
		&i.OpenPrs,
		&i.ClosedPrs,
		&i.MergedPrs,
		&i.GithubCreatedAt,
		&i.GithubUpdatedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getRepositories = `-- name: GetRepositories :many
SELECT DISTINCT
    repositories.id AS repo_id,
    repositories.name AS repo_name,
    organizations.login AS org_login
FROM
    repositories
JOIN
    repository_collaborators ON repositories.id = repository_collaborators.repo_id
JOIN
    organization_collaborators ON repository_collaborators.organization_collaborator_id = organization_collaborators.id
JOIN
    organizations ON organization_collaborators.organization_id = organizations.id ORDER BY repositories.name
`

type GetRepositoriesRow struct {
	RepoID   string         `json:"repo_id"`
	RepoName sql.NullString `json:"repo_name"`
	OrgLogin string         `json:"org_login"`
}

func (q *Queries) GetRepositories(ctx context.Context) ([]GetRepositoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getRepositories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRepositoriesRow
	for rows.Next() {
		var i GetRepositoriesRow
		if err := rows.Scan(&i.RepoID, &i.RepoName, &i.OrgLogin); err != nil {
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

const insertRepo = `-- name: InsertRepo :one
INSERT INTO
    "repositories" (
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
    ) RETURNING repositories.id
`

type InsertRepoParams struct {
	ID              string         `json:"id"`
	Name            sql.NullString `json:"name"`
	IsPrivate       sql.NullBool   `json:"is_private"`
	DefaultBranch   sql.NullString `json:"default_branch"`
	Url             sql.NullString `json:"url"`
	HomepageUrl     sql.NullString `json:"homepage_url"`
	OpenIssues      sql.NullInt32  `json:"open_issues"`
	ClosedIssues    sql.NullInt32  `json:"closed_issues"`
	OpenPrs         sql.NullInt32  `json:"open_prs"`
	ClosedPrs       sql.NullInt32  `json:"closed_prs"`
	MergedPrs       sql.NullInt32  `json:"merged_prs"`
	GithubCreatedAt sql.NullTime   `json:"github_created_at"`
	GithubUpdatedAt sql.NullTime   `json:"github_updated_at"`
}

func (q *Queries) InsertRepo(ctx context.Context, arg InsertRepoParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertRepo,
		arg.ID,
		arg.Name,
		arg.IsPrivate,
		arg.DefaultBranch,
		arg.Url,
		arg.HomepageUrl,
		arg.OpenIssues,
		arg.ClosedIssues,
		arg.OpenPrs,
		arg.ClosedPrs,
		arg.MergedPrs,
		arg.GithubCreatedAt,
		arg.GithubUpdatedAt,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}
