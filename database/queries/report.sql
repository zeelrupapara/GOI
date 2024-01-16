-- name: GetCollaborators :many
select * from collaborators;

-- name: GetCollaboratorByLogin :one
select * from collaborators where login = $1;


-- name: GetCollaboratorsOrgsByCollaboratorID :many
select *
from
    organization_collaborators
where collaborator_id = $1;

-- name: GetRepoCollaboratorsByOrgCollaboratorID :many
select *
from
    repository_collaborators
where
    organization_collaborator_id = $1;

