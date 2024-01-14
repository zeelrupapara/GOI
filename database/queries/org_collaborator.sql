-- name: GetOrgMemberByID :one
SELECT organization_collaborators.id
FROM "organization_collaborators"
WHERE organization_id = $1
    AND collaborator_id = $2;
