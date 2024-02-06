-- name: InsertCommit :one
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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING commits.id;

-- name: GetCommitByID :one
SELECT commits.id FROM "commits" WHERE commits.hash_id = $1 and commits.branch_id = $2;

-- name: GetUserWiseCommitContributionCount :many
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
    CoreData cd ON ds.commit_date = cd.commit_date AND cd.username = ds.username;
