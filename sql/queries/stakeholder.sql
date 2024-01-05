-- LeaderPart

-- name: CreateLeader :one
INSERT INTO leaders (name, mobilephone, emailaddress) VALUES ($1, $2, $3) RETURNING leaderid;

-- name: DeleteLeader :exec
DELETE FROM leaders WHERE leaderid = $1;

-- name: GetLeader :one
SELECT * FROM leaders WHERE leaderid = $1;

-- name: IsLeaderExists :one
SELECT EXISTS(SELECT 1 FROM leaders WHERE name = $1 AND mobilephone = $2 AND emailaddress = $3) AS exists;

-- name: GetLeaderIdByInfo :one
SELECT leaderid FROM leaders WHERE name = $1 AND mobilephone = $2 AND emailaddress = $3;

-- QMPart

-- name: CreateQM :one
INSERT INTO qualitymonitors (name, address, leaderid, contactname, contactphone) VALUES ($1, $2, $3, $4, $5) RETURNING monitorid;

-- name: DeleteQM :exec
DELETE FROM qualitymonitors WHERE monitorid = $1;

-- name: GetQMById :one
SELECT * FROM qualitymonitors WHERE monitorid = $1;

-- name: ListQM :many
SELECT * FROM qualitymonitors;



-- PartnerPart

-- name: CreatePartner :one
INSERT INTO partners (name, address, officephone, leaderid, contactname, contactphone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING partnerid;

-- name: DeletePartner :exec
DELETE FROM partners WHERE partnerid = $1;

-- name: GetPartner :one
SELECT * FROM partners WHERE partnerid = $1;

-- name: ListPartner :many
SELECT * FROM partners;



-- ClientPart

-- name: CreateClient :one
INSERT INTO clients (name, address, officephone, leaderid, contactname, contactphone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING clientid;

-- name: UpdateClient :one
UPDATE clients SET name = $1, address = $2, officephone = $3, leaderid = $4, contactname = $5, contactphone = $6 WHERE clientid = $7 RETURNING *;

-- name: DeleteClient :exec
DELETE FROM clients WHERE clientid = $1;

-- name: GetClient :one
SELECT * FROM clients WHERE clientid = $1;

-- name: ListClient :one
SELECT * FROM clients;


