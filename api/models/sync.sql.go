// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sync.sql

package models

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const getSyncDates = `-- name: GetSyncDates :many
SELECT unique_date,
       ARRAY_AGG(DISTINCT activity_type)::text[] AS activities
FROM (
    SELECT DATE(github_updated_at) AS unique_date, 'Pull Request' AS activity_type
    FROM pull_requests pr 
    UNION ALL
    SELECT DATE(github_committed_time) AS unique_date, 'Commit' AS activity_type
    FROM commits c 
    UNION ALL
    SELECT DATE(github_updated_at) AS unique_date, 'Issue' AS activity_type
    FROM issues i 
) AS combined_dates
GROUP BY unique_date
ORDER BY unique_date
`

type GetSyncDatesRow struct {
	UniqueDate time.Time `json:"unique_date"`
	Activities []string  `json:"activities"`
}

func (q *Queries) GetSyncDates(ctx context.Context) ([]GetSyncDatesRow, error) {
	rows, err := q.db.QueryContext(ctx, getSyncDates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSyncDatesRow
	for rows.Next() {
		var i GetSyncDatesRow
		if err := rows.Scan(&i.UniqueDate, pq.Array(&i.Activities)); err != nil {
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
