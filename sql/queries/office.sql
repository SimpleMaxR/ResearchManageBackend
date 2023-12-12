-- name: CreateOffice :one
INSERT INTO offices (lab_id, area, address, managerid) VALUES ($1, $2, $3, $4)
RETURNING officeid;

-- name: UpdateOffice :one
UPDATE offices SET lab_id = $2, area = $3, address = $4, managerid = $5 WHERE officeid = $1
RETURNING *;

-- name: DeleteOffice :exec
DELETE FROM offices WHERE officeid = $1;

-- name: ListOfficeAll :many
SELECT * FROM offices;

-- name: GetOfficeByLab :one
SELECT * FROM offices WHERE lab_id = $1;
