-- name: InsertRepo :one
INSERT INTO
    "repositories" (
        "id",
        "name",
        "is_private",
        "default_branch",
        "url",
        "homepage_url",
        "open_issues",
        "closed_issues",
        "open_prs",
        "closed_prs",
        "merged_prs",
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
    ) RETURNING repositories.id;
-- name: GetRepoByID :one
SELECT repositories.id
FROM "repositories"
WHERE repositories.id = $1;

-- name: GetRepoDetailsByID :one
select * from repositories where id = $1;

-- name: GetRepositories :many
SELECT DISTINCT
    repositories.id AS repo_id,
    repositories.name AS repo_name,
    organizations.login AS org_login
FROM
    repositories
JOIN
    repository_collaborators ON repositories.id = repository_collaborators.repo_id
JOIN
    organization_collaborators ON repository_collaborators.organization_collaborator_id = organization_collaborators.id
JOIN
    organizations ON organization_collaborators.organization_id = organizations.id ORDER BY repositories.name;

-- name: GetRepoCountByFilters :one
SELECT 
	COUNT(DISTINCT  r.id)
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
    AND rc.repo_id = ANY(string_to_array($5, ','));

-- name: GetRepoIDs :many
SELECT DISTINCT
    repositories.id
FROM
    "repositories";
