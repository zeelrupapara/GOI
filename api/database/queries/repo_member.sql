-- name: GetRepoMemberByOrgRepoID :one
SELECT
    repository_collaborators.id
FROM
    "repository_collaborators"
WHERE
    repository_collaborators.repo_id = $1
    AND repository_collaborators.organization_collaborator_id = $2;

-- name: InsertOrgRepoMember :one
INSERT INTO
    "repository_collaborators" (
        "id",
        "repo_id",
        "organization_collaborator_id"
    )
VALUES ($1, $2, $3) RETURNING repository_collaborators.id;
