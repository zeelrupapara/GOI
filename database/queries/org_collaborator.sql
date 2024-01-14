-- name: GetOrgMemberByID :one
SELECT organization_collaborators.id
FROM "organization_collaborators"
WHERE organization_id = $1
    AND collaborator_id = $2;
-- name: InsertOrgMember :one
INSERT INTO "organization_collaborators" (
        "organization_id",
        "collaborator_id"
    )
VALUES ($1, $2)
RETURNING organization_collaborators.id;
