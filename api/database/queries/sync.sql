-- name: GetSyncDates :many
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
ORDER BY unique_date;
