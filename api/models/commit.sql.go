// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: commit.sql

package models

import (
	"context"
	"database/sql"
	"time"
)

const getCommitByID = `-- name: GetCommitByID :one
SELECT commits.id FROM "commits" WHERE commits.hash_id = $1 and commits.branch_id = $2
`

type GetCommitByIDParams struct {
	HashID   string `json:"hash_id"`
	BranchID string `json:"branch_id"`
}

func (q *Queries) GetCommitByID(ctx context.Context, arg GetCommitByIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getCommitByID, arg.HashID, arg.BranchID)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getDefaultBranchCommitByFilters = `-- name: GetDefaultBranchCommitByFilters :many
SELECT distinct 
    coll.login as commiter,
    r.name as repository,
    o.login as organization,
    c.message as message,
    c.github_committed_time as commit_date
FROM commits c 
JOIN branches b on b.id = c.branch_id 
JOIN repositories r on r.id = b.repository_id
JOIN repository_collaborators rc on rc.repo_id = r.id 
JOIN organization_collaborators oc on oc.id = rc.organization_collaborator_id 
JOIN organizations o on o.id = oc.organization_id 
JOIN collaborators coll on coll.id = c.author_id  
WHERE (c.github_committed_time between $1 AND $2)
    AND b.is_default = true
    AND coll.login = $3
    AND o.login = $4
    AND r.name = $5
ORDER BY commit_date DESC
`

type GetDefaultBranchCommitByFiltersParams struct {
	GithubCommittedTime   sql.NullTime   `json:"github_committed_time"`
	GithubCommittedTime_2 sql.NullTime   `json:"github_committed_time_2"`
	Login                 string         `json:"login"`
	Login_2               string         `json:"login_2"`
	Name                  sql.NullString `json:"name"`
}

type GetDefaultBranchCommitByFiltersRow struct {
	Commiter     string         `json:"commiter"`
	Repository   sql.NullString `json:"repository"`
	Organization string         `json:"organization"`
	Message      sql.NullString `json:"message"`
	CommitDate   sql.NullTime   `json:"commit_date"`
}

func (q *Queries) GetDefaultBranchCommitByFilters(ctx context.Context, arg GetDefaultBranchCommitByFiltersParams) ([]GetDefaultBranchCommitByFiltersRow, error) {
	rows, err := q.db.QueryContext(ctx, getDefaultBranchCommitByFilters,
		arg.GithubCommittedTime,
		arg.GithubCommittedTime_2,
		arg.Login,
		arg.Login_2,
		arg.Name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDefaultBranchCommitByFiltersRow
	for rows.Next() {
		var i GetDefaultBranchCommitByFiltersRow
		if err := rows.Scan(
			&i.Commiter,
			&i.Repository,
			&i.Organization,
			&i.Message,
			&i.CommitDate,
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

const getRepoWiseCommitContributionDetailsByFilters = `-- name: GetRepoWiseCommitContributionDetailsByFilters :many
SELECT distinct 
    coll.login as commiter,
    r.name as repository,
    b.name as branch,
    o.login as organization,
    count(distinct c.id) as commits,
    date(c.github_committed_time) as commit_date
FROM commits c 
JOIN branches b on b.id = c.branch_id 
JOIN repositories r on r.id = b.repository_id
JOIN repository_collaborators rc on rc.repo_id = r.id 
JOIN organization_collaborators oc on oc.id = rc.organization_collaborator_id 
JOIN organizations o on o.id = oc.organization_id 
JOIN collaborators coll on coll.id = c.author_id  
WHERE (c.github_committed_time between $1 and $2)
    AND b.is_default = true
    AND coll.id = ANY(string_to_array($3, ','))
    AND o.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','))
GROUP BY coll.login, date(c.github_committed_time), r.name, b.name ,o.login
ORDER BY commit_date DESC LIMIT $6 OFFSET $7
`

type GetRepoWiseCommitContributionDetailsByFiltersParams struct {
	GithubCommittedTime   sql.NullTime `json:"github_committed_time"`
	GithubCommittedTime_2 sql.NullTime `json:"github_committed_time_2"`
	StringToArray         string       `json:"string_to_array"`
	StringToArray_2       string       `json:"string_to_array_2"`
	StringToArray_3       string       `json:"string_to_array_3"`
	Limit                 int32        `json:"limit"`
	Offset                int32        `json:"offset"`
}

type GetRepoWiseCommitContributionDetailsByFiltersRow struct {
	Commiter     string         `json:"commiter"`
	Repository   sql.NullString `json:"repository"`
	Branch       string         `json:"branch"`
	Organization string         `json:"organization"`
	Commits      int64          `json:"commits"`
	CommitDate   time.Time      `json:"commit_date"`
}

func (q *Queries) GetRepoWiseCommitContributionDetailsByFilters(ctx context.Context, arg GetRepoWiseCommitContributionDetailsByFiltersParams) ([]GetRepoWiseCommitContributionDetailsByFiltersRow, error) {
	rows, err := q.db.QueryContext(ctx, getRepoWiseCommitContributionDetailsByFilters,
		arg.GithubCommittedTime,
		arg.GithubCommittedTime_2,
		arg.StringToArray,
		arg.StringToArray_2,
		arg.StringToArray_3,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRepoWiseCommitContributionDetailsByFiltersRow
	for rows.Next() {
		var i GetRepoWiseCommitContributionDetailsByFiltersRow
		if err := rows.Scan(
			&i.Commiter,
			&i.Repository,
			&i.Branch,
			&i.Organization,
			&i.Commits,
			&i.CommitDate,
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

const getUserWiseCommitContributionCount = `-- name: GetUserWiseCommitContributionCount :many
WITH CoreData AS (
SELECT distinct 
    coll.login as username,
    count(distinct c.id) as total_commit,
    date(c.github_committed_time) as commit_date
FROM commits c 
JOIN branches b on b.id = c.branch_id 
JOIN repositories r on r.id = b.repository_id
JOIN repository_collaborators rc on rc.repo_id = r.id 
JOIN organization_collaborators oc on oc.id = rc.organization_collaborator_id 
JOIN organizations o on o.id = oc.organization_id 
JOIN collaborators coll on coll.id = c.author_id  
WHERE (c.github_committed_time between $1 and $2)
    AND b.is_default = true
    AND coll.id = ANY(string_to_array($3, ','))
    AND o.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','))
GROUP BY coll.login, date(c.github_committed_time)),
DateSeries AS (
    SELECT commit_date::date, username
    FROM (
        SELECT generate_series((SELECT min(commit_date) - interval '1 day' FROM CoreData), 
                               (SELECT max(commit_date) + interval '1 day' FROM CoreData), 
                               interval '1 day') AS commit_date
    ) x
    CROSS JOIN (
        SELECT DISTINCT username
        FROM CoreData
        WHERE username IS NOT NULL
    ) y
)
SELECT 
    ds.commit_date, 
    ds.username, 
    COALESCE(cd.total_commit, 0) AS total_commit 
FROM 
    DateSeries ds 
LEFT JOIN 
    CoreData cd ON ds.commit_date = cd.commit_date AND cd.username = ds.username
`

type GetUserWiseCommitContributionCountParams struct {
	GithubCommittedTime   sql.NullTime `json:"github_committed_time"`
	GithubCommittedTime_2 sql.NullTime `json:"github_committed_time_2"`
	StringToArray         string       `json:"string_to_array"`
	StringToArray_2       string       `json:"string_to_array_2"`
	StringToArray_3       string       `json:"string_to_array_3"`
}

type GetUserWiseCommitContributionCountRow struct {
	CommitDate  time.Time `json:"commit_date"`
	Username    string    `json:"username"`
	TotalCommit int64     `json:"total_commit"`
}

func (q *Queries) GetUserWiseCommitContributionCount(ctx context.Context, arg GetUserWiseCommitContributionCountParams) ([]GetUserWiseCommitContributionCountRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserWiseCommitContributionCount,
		arg.GithubCommittedTime,
		arg.GithubCommittedTime_2,
		arg.StringToArray,
		arg.StringToArray_2,
		arg.StringToArray_3,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserWiseCommitContributionCountRow
	for rows.Next() {
		var i GetUserWiseCommitContributionCountRow
		if err := rows.Scan(&i.CommitDate, &i.Username, &i.TotalCommit); err != nil {
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

const insertCommit = `-- name: InsertCommit :one
INSERT INTO
    "commits" (
        "id",
        "hash_id",
        "message",
        "branch_id",
        "author_id",
        "pr_id",
        "url",
        "commit_url",
        "github_committed_time"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING commits.id
`

type InsertCommitParams struct {
	ID                  string         `json:"id"`
	HashID              string         `json:"hash_id"`
	Message             sql.NullString `json:"message"`
	BranchID            string         `json:"branch_id"`
	AuthorID            string         `json:"author_id"`
	PrID                sql.NullString `json:"pr_id"`
	Url                 sql.NullString `json:"url"`
	CommitUrl           sql.NullString `json:"commit_url"`
	GithubCommittedTime sql.NullTime   `json:"github_committed_time"`
}

func (q *Queries) InsertCommit(ctx context.Context, arg InsertCommitParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertCommit,
		arg.ID,
		arg.HashID,
		arg.Message,
		arg.BranchID,
		arg.AuthorID,
		arg.PrID,
		arg.Url,
		arg.CommitUrl,
		arg.GithubCommittedTime,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}
