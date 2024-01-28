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

