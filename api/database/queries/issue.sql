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
    public.issues i
JOIN
    public.repository_collaborators rc ON i.repository_collaborators_id = rc.id
JOIN
    public.repositories r ON rc.repo_id = r.id
JOIN
    public.organization_collaborators oc ON rc.organization_collaborator_id = oc.id
JOIN
    public.organizations org ON oc.organization_id = org.id
JOIN
    public.assignees a ON i.id = a.issue_id
WHERE
    i.github_updated_at BETWEEN $1 AND $2
    AND a.collaborator_id = ANY(string_to_array($3, ','))
    AND org.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','));

