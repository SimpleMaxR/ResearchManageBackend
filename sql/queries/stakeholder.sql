-- ContactsPart

-- name: CreateContact :one
INSERT INTO contacts (name, officephone, mobilephone, emailaddress) VALUES ($1, $2, $3, $4) RETURNING contactid;

-- name: SetContactClient :one
UPDATE contacts SET baseclient = $1 WHERE contactid = $2 RETURNING contactid;

-- name: SetContactPartner :one
UPDATE contacts SET basepartners = $1 WHERE contactid = $2 RETURNING contactid;

-- name: SetContactQM :one
UPDATE contacts SET baseqm = $1 WHERE contactid = $2 RETURNING contactid;



-- LeaderPart

-- name: CreateLeader :one
INSERT INTO leaders (name, officephone, mobilephone, emailaddress) VALUES ($1, $2, $3, $4) RETURNING leaderid;

-- name: DeleteLeader :exec
DELETE FROM leaders WHERE leaderid = $1;

-- name: GetLeader :one
SELECT * FROM leaders WHERE leaderid = $1;



-- QMPart

-- name: CreateQM :one
INSERT INTO qualitymonitors (name, address, leaderid) VALUES ($1, $2, $3) RETURNING monitorid;

-- name: DeleteQM :exec
DELETE FROM qualitymonitors WHERE monitorid = $1;

-- name: GetQMById :one
SELECT * FROM qualitymonitors WHERE monitorid = $1;

-- name: ListQM :one
SELECT * FROM qualitymonitors;



-- PartnerPart

-- name: CreatePartner :one
INSERT INTO partners (name, address, officephone, leaderid) VALUES ($1, $2, $3, $4) RETURNING partnerid;

-- name: DeletePartner :exec
DELETE FROM partners WHERE partnerid = $1;

-- name: GetPartner :one
SELECT * FROM partners WHERE partnerid = $1;

-- name: ListPartner :many
SELECT * FROM partners;



-- ClientPart

-- name: CreateClient :one
INSERT INTO clients (name, address, officephone, leaderid) VALUES ($1, $2, $3, $4) RETURNING clientid;

-- name: UpdateClient :one
UPDATE clients SET name = $1, address = $2, officephone = $3, leaderid = $4 WHERE clientid = $5 RETURNING *;

-- name: DeleteClient :exec
DELETE FROM clients WHERE clientid = $1;

-- name: GetClient :one
SELECT * FROM clients WHERE clientid = $1;

-- name: ListClient :one
SELECT * FROM clients;


