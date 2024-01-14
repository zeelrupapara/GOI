-- name: InsertAssignedLabal :one
INSERT INTO
    "assigned_labals" (
        "id",
        "labal_id",
        "activity_id",
        "activity_type"
    )
VALUES ($1, $2, $3, $4) RETURNING assigned_labals.id;

-- name: GetAssignedLabal :one
SELECT assigned_labals.id
FROM "assigned_labals"
WHERE assigned_labals.labal_id = $1 AND assigned_labals.activity_id = $2;

-- name: InsertAssignee :one
INSERT INTO
    "assignees" (
        "id",
        "collaborator_id",
        "activity_id",
        "activity_type"
    )
VALUES ($1, $2, $3, $4) RETURNING assignees.id;

-- name: GetAssigneeByID :one
SELECT assignees.id FROM "assignees" WHERE assignees.collaborator_id = $1 AND assignees.activity_id = $2;
