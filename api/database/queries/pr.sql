-- name: InsertPR :one
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
    ) RETURNING pull_requests.id;
-- name: GetPRByID :one
SELECT pull_requests.id
FROM "pull_requests"
WHERE pull_requests.id = $1;

-- name: GetPRCountByFilters :one
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
    AND r.id = ANY(string_to_array($5, ','));

-- name: GetPullRequestContributionByFilters :many
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
GROUP BY DATE(pr.github_updated_at);

-- name: GetPullRequestContributionDetailsByFilters :many
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
ORDER BY pr.github_updated_at DESC LIMIT $6 OFFSET $7;

-- name: GetUserWisePullRequestContributionByFilters :many
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
    CoreData cd ON ds.user_date = cd.updated_date AND cd.login = ds.login;
