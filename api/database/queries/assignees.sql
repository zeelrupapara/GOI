-- name: InsertAssignedLabal :one
INSERT INTO
    "assigned_labals" (
        "id",
        "labal_id",
        "pr_id",
        "issue_id",
        "activity_type"
    )
VALUES ($1, $2, $3, $4, $5) RETURNING assigned_labals.id;

-- name: GetAssignedLabalByPR :one
SELECT assigned_labals.id
FROM "assigned_labals"
WHERE
    assigned_labals.labal_id = $1
    AND assigned_labals.pr_id = $2;

-- name: GetAssignedLabalByIssue :one
SELECT assigned_labals.id
FROM "assigned_labals"
WHERE
    assigned_labals.labal_id = $1
    AND assigned_labals.issue_id = $2;

-- name: InsertAssignee :one
INSERT INTO
    "assignees" (
        "id",
        "collaborator_id",
        "pr_id",
        "issue_id",
        "activity_type"
    )
VALUES ($1, $2, $3, $4, $5) RETURNING assignees.id;

-- name: GetAssigneeByPR :one
SELECT assignees.id
FROM "assignees"
WHERE
    assignees.collaborator_id = $1
    AND assignees.pr_id = $2;

-- name: GetAssigneeByIssue :one
SELECT assignees.id
FROM "assignees"
WHERE
    assignees.collaborator_id = $1
    AND assignees.issue_id = $2;
