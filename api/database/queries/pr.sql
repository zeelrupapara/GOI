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
    public.pull_requests pr
JOIN
    public.repository_collaborators rc ON pr.repository_collaborators_id = rc.id
JOIN
    public.repositories r ON rc.repo_id = r.id
JOIN
    public.organization_collaborators oc ON rc.organization_collaborator_id = oc.id
JOIN
    public.organizations org ON oc.organization_id = org.id
JOIN
    public.assignees a ON pr.id = a.pr_id
LEFT JOIN
    public.collaborators coll ON a.collaborator_id = coll.id
WHERE
    pr.github_updated_at BETWEEN $1 AND $2 
    AND coll.id = ANY(string_to_array($3, ','))
    AND org.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','));
