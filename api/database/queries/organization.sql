-- name: GetOrganizations :many
SELECT *
FROM "organizations"
ORDER BY organizations.login;

-- name: InsertOrganization :one
INSERT INTO "organizations" (
        "id",
        "login",
        "name",
        "email",
        "location",
        "description",
        "url",
        "avatar_url",
        "website_url",
        "github_created_at",
        "github_updated_at"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING organizations.id;

-- name: GetOrganizationByLogin :one
SELECT organizations.id
FROM "organizations"
WHERE login = $1;

-- name: GetOrganizationByFilter :one
SELECT
    COUNT(DISTINCT org.id) AS organization_count
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

-- name: GetOrganizationIDs :many
SELECT DISTINCT
    organizations.id
FROM
    "organizations";

-- name: GetOrganizationContributionsByFilters :many
SELECT DISTINCT
    org.login AS organization_name,
    COUNT(DISTINCT pr.id) AS total_prs,
    COUNT(DISTINCT i.id) AS total_issues
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
GROUP BY org.login;
