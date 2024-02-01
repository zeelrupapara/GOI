-- name: InsertIssue :one
INSERT INTO
    "issues" (
        "id",
        "title",
        "status",
        "url",
        "number",
        "author_id",
        "repository_collaborators_id",
        "github_closed_at",
        "github_created_at",
        "github_updated_at"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING issues.id;

-- name: GetIssueByID :one
SELECT issues.id FROM "issues" WHERE issues.id = $1;

-- name: GetIssueCountByFilters :one
SELECT
    COUNT(DISTINCT i.id) AS issue_count
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
    AND a.collaborator_id = ANY(string_to_array($3, ','))
    AND org.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','));

-- name: GetIssueContributionByFilters :many
SELECT
    DATE(i.github_updated_at),
    COUNT(DISTINCT i.id) FILTER(WHERE i.status = 'OPEN') AS total_open_issues,
    COUNT(DISTINCT i.id) FILTER(WHERE i.status = 'CLOSED') AS total_closed_issues
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
    (i.github_updated_at BETWEEN $1 AND $2)   
    AND coll.id = ANY(string_to_array($3, ','))
    AND org.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','))
GROUP BY DATE(i.github_updated_at);

-- name: GetIssueContributionDetailsByFilters :many
SELECT DISTINCT
    i.id AS id,
    i.url AS url,
    i.title AS title,
    i.status AS status,
    coll.login AS assignee_name,
    r.name AS repository_name,
    org.login AS organization_name,
    i.github_updated_at AS updated_at
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
    (i.github_updated_at BETWEEN $1 AND $2)   
    AND coll.id = ANY(string_to_array($3, ','))
    AND org.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','))
ORDER BY i.github_updated_at DESC LIMIT $6 OFFSET $7;

-- name: GetUserWiseIssueContributionByFilters :many
WITH CoreData AS (
    SELECT
        COUNT(DISTINCT i.id) AS issue_count,
        DATE(i.github_updated_at) AS updated_date,
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
        (i.github_updated_at BETWEEN $1 AND $2)   
        AND coll.id = ANY(string_to_array($3, ','))
        AND org.id = ANY(string_to_array($4, ','))
        AND r.id = ANY(string_to_array($5, ','))
        AND i.status = $6
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
    COALESCE(cd.issue_count, 0) AS issue_count 
FROM 
    DateSeries ds 
LEFT JOIN 
    CoreData cd ON ds.user_date = cd.updated_date AND cd.login = ds.login;
