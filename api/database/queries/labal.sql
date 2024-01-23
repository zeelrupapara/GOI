-- name: InsertLabal :one
INSERT INTO
    "labals" ("id", "name")
VALUES ($1, $2) RETURNING labals.id;

-- name: GetLabalByID :one
SELECT labals.id FROM "labals" WHERE labals.id = $1;
