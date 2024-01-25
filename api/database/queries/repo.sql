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
