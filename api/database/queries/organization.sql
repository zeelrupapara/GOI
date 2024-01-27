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
    public.organizations org
FULL JOIN
    public.organization_collaborators oc ON org.id = oc.organization_id
FULL JOIN
    public.collaborators coll ON oc.collaborator_id = coll.id
FULL JOIN
    public.repository_collaborators rc ON oc.id = rc.organization_collaborator_id
FULL JOIN
    public.repositories r ON rc.repo_id = r.id
FULL JOIN
    public.issues i ON rc.id = i.repository_collaborators_id
FULL JOIN
    public.pull_requests pr ON rc.id = pr.repository_collaborators_id
FULL JOIN
    public.assignees a ON (i.id = a.issue_id OR pr.id = a.pr_id)
WHERE
	(i.github_updated_at BETWEEN $1 AND $2 OR pr.github_updated_at BETWEEN $1 AND $2)
    AND coll.id = ANY(string_to_array($3, ','))
    AND org.id = ANY(string_to_array($4, ','))
    AND r.id = ANY(string_to_array($5, ','));

-- name: GetOrganizationIDs :many
SELECT DISTINCT
    organizations.id
FROM
    "organizations";
