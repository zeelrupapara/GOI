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
