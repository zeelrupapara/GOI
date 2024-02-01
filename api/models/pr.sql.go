// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: pr.sql

package models

import (
	"context"
	"database/sql"
	"time"
)

const getPRByID = `-- name: GetPRByID :one
SELECT pull_requests.id
FROM "pull_requests"
WHERE pull_requests.id = $1
`

func (q *Queries) GetPRByID(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, getPRByID, id)
	err := row.Scan(&id)
	return id, err
}

const getPRCountByFilters = `-- name: GetPRCountByFilters :one
SELECT
   	COUNT(DISTINCT pr.id)
FROM
    public.repositories r
JOIN
    public.repository_collaborators rc ON r.id = rc.repo_id
JOIN
    public.organization_collaborators oc ON rc.organization_collaborator_id = oc.id
JOIN
    public.organizations org ON oc.organization_id = org.id
LEFT JOIN
    public.issues i ON rc.id = i.repository_collaborators_id
LEFT JOIN
    public.pull_requests pr ON rc.id = pr.repository_collaborators_id
LEFT JOIN
    public.assignees a ON (i.id = a.issue_id OR pr.id = a.pr_id)
LEFT JOIN
    public.collaborators coll ON a.collaborator_id = coll.id
WHERE
    (
        (pr.github_updated_at BETWEEN $1 AND $2) OR
        (i.github_updated_at BETWEEN $1 AND $2)
    )
    AND coll.id = ANY(string_to_array($3, ','))
    AND org.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','))
`

type GetPRCountByFiltersParams struct {
	GithubUpdatedAt   sql.NullTime `json:"github_updated_at"`
	GithubUpdatedAt_2 sql.NullTime `json:"github_updated_at_2"`
	StringToArray     string       `json:"string_to_array"`
	StringToArray_2   string       `json:"string_to_array_2"`
	StringToArray_3   string       `json:"string_to_array_3"`
}

func (q *Queries) GetPRCountByFilters(ctx context.Context, arg GetPRCountByFiltersParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getPRCountByFilters,
		arg.GithubUpdatedAt,
		arg.GithubUpdatedAt_2,
		arg.StringToArray,
		arg.StringToArray_2,
		arg.StringToArray_3,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getPullRequestContributionByFilters = `-- name: GetPullRequestContributionByFilters :many
SELECT
    DATE(pr.github_updated_at),
    COUNT(DISTINCT pr.id) FILTER(WHERE pr.status = 'OPEN') AS total_open_prs,
    COUNT(DISTINCT pr.id) FILTER(WHERE pr.status = 'CLOSED') AS total_closed_prs,
    COUNT(DISTINCT pr.id) FILTER(WHERE pr.status = 'MERGED') AS total_merged_prs
FROM
    public.repositories r
JOIN
    public.repository_collaborators rc ON r.id = rc.repo_id
JOIN
    public.organization_collaborators oc ON rc.organization_collaborator_id = oc.id
JOIN
    public.organizations org ON oc.organization_id = org.id
LEFT JOIN
    public.issues i ON rc.id = i.repository_collaborators_id
LEFT JOIN
    public.pull_requests pr ON rc.id = pr.repository_collaborators_id
LEFT JOIN
    public.assignees a ON (i.id = a.issue_id OR pr.id = a.pr_id)
LEFT JOIN
    public.collaborators coll ON a.collaborator_id = coll.id
WHERE
    (pr.github_updated_at BETWEEN $1 AND $2)   
    AND coll.id = ANY(string_to_array($3, ','))
    AND org.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','))
GROUP BY DATE(pr.github_updated_at)
`

type GetPullRequestContributionByFiltersParams struct {
	GithubUpdatedAt   sql.NullTime `json:"github_updated_at"`
	GithubUpdatedAt_2 sql.NullTime `json:"github_updated_at_2"`
	StringToArray     string       `json:"string_to_array"`
	StringToArray_2   string       `json:"string_to_array_2"`
	StringToArray_3   string       `json:"string_to_array_3"`
}

type GetPullRequestContributionByFiltersRow struct {
	Date           time.Time `json:"date"`
	TotalOpenPrs   int64     `json:"total_open_prs"`
	TotalClosedPrs int64     `json:"total_closed_prs"`
	TotalMergedPrs int64     `json:"total_merged_prs"`
}

func (q *Queries) GetPullRequestContributionByFilters(ctx context.Context, arg GetPullRequestContributionByFiltersParams) ([]GetPullRequestContributionByFiltersRow, error) {
	rows, err := q.db.QueryContext(ctx, getPullRequestContributionByFilters,
		arg.GithubUpdatedAt,
		arg.GithubUpdatedAt_2,
		arg.StringToArray,
		arg.StringToArray_2,
		arg.StringToArray_3,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPullRequestContributionByFiltersRow
	for rows.Next() {
		var i GetPullRequestContributionByFiltersRow
		if err := rows.Scan(
			&i.Date,
			&i.TotalOpenPrs,
			&i.TotalClosedPrs,
			&i.TotalMergedPrs,
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

const getPullRequestContributionDetailsByFilters = `-- name: GetPullRequestContributionDetailsByFilters :many
SELECT DISTINCT
    pr.id AS id,
    pr.url AS url,
    pr.title AS title,
    pr.status AS status,
    coll.login AS assignee_name,
    r.name AS repository_name,
    org.login AS organization_name,
    pr.github_updated_at AS updated_at
FROM
    public.repositories r
JOIN
    public.repository_collaborators rc ON r.id = rc.repo_id
JOIN
    public.organization_collaborators oc ON rc.organization_collaborator_id = oc.id
JOIN
    public.organizations org ON oc.organization_id = org.id
LEFT JOIN
    public.issues i ON rc.id = i.repository_collaborators_id
LEFT JOIN
    public.pull_requests pr ON rc.id = pr.repository_collaborators_id
LEFT JOIN
    public.assignees a ON (i.id = a.issue_id OR pr.id = a.pr_id)
LEFT JOIN
    public.collaborators coll ON a.collaborator_id = coll.id
WHERE
    (pr.github_updated_at BETWEEN $1 AND $2)   
    AND coll.id = ANY(string_to_array($3, ','))
    AND org.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','))
    AND pr.status = ANY(string_to_array($8, ','))
ORDER BY pr.github_updated_at DESC LIMIT $6 OFFSET $7
`

type GetPullRequestContributionDetailsByFiltersParams struct {
	GithubUpdatedAt   sql.NullTime `json:"github_updated_at"`
	GithubUpdatedAt_2 sql.NullTime `json:"github_updated_at_2"`
	StringToArray     string       `json:"string_to_array"`
	StringToArray_2   string       `json:"string_to_array_2"`
	StringToArray_3   string       `json:"string_to_array_3"`
	Limit             int32        `json:"limit"`
	Offset            int32        `json:"offset"`
	StringToArray_4   string       `json:"string_to_array_4"`
}

type GetPullRequestContributionDetailsByFiltersRow struct {
	ID               sql.NullString `json:"id"`
	Url              sql.NullString `json:"url"`
	Title            sql.NullString `json:"title"`
	Status           sql.NullString `json:"status"`
	AssigneeName     sql.NullString `json:"assignee_name"`
	RepositoryName   sql.NullString `json:"repository_name"`
	OrganizationName string         `json:"organization_name"`
	UpdatedAt        sql.NullTime   `json:"updated_at"`
}

func (q *Queries) GetPullRequestContributionDetailsByFilters(ctx context.Context, arg GetPullRequestContributionDetailsByFiltersParams) ([]GetPullRequestContributionDetailsByFiltersRow, error) {
	rows, err := q.db.QueryContext(ctx, getPullRequestContributionDetailsByFilters,
		arg.GithubUpdatedAt,
		arg.GithubUpdatedAt_2,
		arg.StringToArray,
		arg.StringToArray_2,
		arg.StringToArray_3,
		arg.Limit,
		arg.Offset,
		arg.StringToArray_4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPullRequestContributionDetailsByFiltersRow
	for rows.Next() {
		var i GetPullRequestContributionDetailsByFiltersRow
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.Title,
			&i.Status,
			&i.AssigneeName,
			&i.RepositoryName,
			&i.OrganizationName,
			&i.UpdatedAt,
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

const getUserWisePullRequestContributionByFilters = `-- name: GetUserWisePullRequestContributionByFilters :many
WITH CoreData AS (
    SELECT
        COUNT(DISTINCT pr.id) AS pr_count,
        DATE(pr.github_updated_at) AS updated_date,
        coll.login 
    FROM
        public.repositories r
    JOIN
        public.repository_collaborators rc ON r.id = rc.repo_id
    JOIN
        public.organization_collaborators oc ON rc.organization_collaborator_id = oc.id
    JOIN
        public.organizations org ON oc.organization_id = org.id
    LEFT JOIN
        public.issues i ON rc.id = i.repository_collaborators_id
    LEFT JOIN
        public.pull_requests pr ON rc.id = pr.repository_collaborators_id
    LEFT JOIN
        public.assignees a ON (i.id = a.issue_id OR pr.id = a.pr_id)
    LEFT JOIN
        public.collaborators coll ON a.collaborator_id = coll.id
    WHERE
        (pr.github_updated_at BETWEEN $1 AND $2)   
        AND coll.id = ANY(string_to_array($3, ','))
        AND org.id = ANY(string_to_array($4, ','))
        AND r.id = ANY(string_to_array($5, ','))
        AND pr.status = $6
    GROUP BY updated_date, coll.login  
    ORDER BY updated_date DESC
),
DateSeries AS (
    SELECT user_date::date, login
    FROM (
        SELECT generate_series((SELECT min(updated_date) - interval '1 day' FROM CoreData), 
                               (SELECT max(updated_date) + interval '1 day' FROM CoreData), 
                               interval '1 day') AS user_date
    ) x
    CROSS JOIN (
        SELECT DISTINCT login
        FROM CoreData
        WHERE login IS NOT NULL
    ) y
)
SELECT 
    ds.user_date, 
    ds.login, 
    COALESCE(cd.pr_count, 0) AS pr_count 
FROM 
    DateSeries ds 
LEFT JOIN 
    CoreData cd ON ds.user_date = cd.updated_date AND cd.login = ds.login
`

type GetUserWisePullRequestContributionByFiltersParams struct {
	GithubUpdatedAt   sql.NullTime   `json:"github_updated_at"`
	GithubUpdatedAt_2 sql.NullTime   `json:"github_updated_at_2"`
	StringToArray     string         `json:"string_to_array"`
	StringToArray_2   string         `json:"string_to_array_2"`
	StringToArray_3   string         `json:"string_to_array_3"`
	Status            sql.NullString `json:"status"`
}

type GetUserWisePullRequestContributionByFiltersRow struct {
	UserDate time.Time      `json:"user_date"`
	Login    sql.NullString `json:"login"`
	PrCount  int64          `json:"pr_count"`
}

func (q *Queries) GetUserWisePullRequestContributionByFilters(ctx context.Context, arg GetUserWisePullRequestContributionByFiltersParams) ([]GetUserWisePullRequestContributionByFiltersRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserWisePullRequestContributionByFilters,
		arg.GithubUpdatedAt,
		arg.GithubUpdatedAt_2,
		arg.StringToArray,
		arg.StringToArray_2,
		arg.StringToArray_3,
		arg.Status,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserWisePullRequestContributionByFiltersRow
	for rows.Next() {
		var i GetUserWisePullRequestContributionByFiltersRow
		if err := rows.Scan(&i.UserDate, &i.Login, &i.PrCount); err != nil {
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

const insertPR = `-- name: InsertPR :one
INSERT INTO
    "pull_requests" (
        "id",
        "title",
        "status",
        "url",
        "number",
        "is_draft",
        "branch",
        "author_id",
        "repository_collaborators_id",
        "github_closed_at",
        "github_merged_at",
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
    ) RETURNING pull_requests.id
`

type InsertPRParams struct {
	ID                        string         `json:"id"`
	Title                     sql.NullString `json:"title"`
	Status                    sql.NullString `json:"status"`
	Url                       sql.NullString `json:"url"`
	Number                    sql.NullInt32  `json:"number"`
	IsDraft                   sql.NullBool   `json:"is_draft"`
	Branch                    sql.NullString `json:"branch"`
	AuthorID                  string         `json:"author_id"`
	RepositoryCollaboratorsID string         `json:"repository_collaborators_id"`
	GithubClosedAt            sql.NullTime   `json:"github_closed_at"`
	GithubMergedAt            sql.NullTime   `json:"github_merged_at"`
	GithubCreatedAt           sql.NullTime   `json:"github_created_at"`
	GithubUpdatedAt           sql.NullTime   `json:"github_updated_at"`
}

func (q *Queries) InsertPR(ctx context.Context, arg InsertPRParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertPR,
		arg.ID,
		arg.Title,
		arg.Status,
		arg.Url,
		arg.Number,
		arg.IsDraft,
		arg.Branch,
		arg.AuthorID,
		arg.RepositoryCollaboratorsID,
		arg.GithubClosedAt,
		arg.GithubMergedAt,
		arg.GithubCreatedAt,
		arg.GithubUpdatedAt,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}
