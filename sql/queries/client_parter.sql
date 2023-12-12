-- Leader

-- name: CreateLeader :one
INSERT INTO leaders (name, officephone, mobilephone, emailaddress) VALUES ($1, $2, $3, $4) 
RETURNING leaderid;

-- name: DeleteLeader :exec
DELETE FROM leaders WHERE leaderid = $1;

-- name: ListLeaderAll :many
SELECT * FROM leaders;

-- name: ListLeader :one
SELECT * FROM leaders WHERE leaderid = $1;



-- clients

-- name: CreateClient :one
INSERT INTO clients (name, address, leaderid, officephone) VALUES ($1, $2, $3, $4) 
RETURNING clientid;

-- name: DeleteClient :exec
DELETE FROM clients WHERE clientid = $1;

-- name: ListClientAll :many
SELECT * FROM clients;


-- partners

-- name: CreatePartner :one
INSERT INTO partners (name, address, leaderid, officephone) VALUES ($1, $2, $3, $4) 
RETURNING partnerid;

-- name: DeletePartner :exec
DELETE FROM partners WHERE partnerid = $1;

-- name: ListPartnerAll :many
SELECT * FROM partners;


-- qualitymonitors

-- name: CreateQualitymonitor :one
INSERT INTO qualitymonitors (name, address, leaderid) VALUES ($1, $2, $3) 
RETURNING monitorid;

-- name: DeleteQualitymonitor :exec
DELETE FROM qualitymonitors WHERE monitorid = $1;

-- name: ListQualitymonitorAll :many
SELECT * FROM qualitymonitors;


-- contacts

-- name: CreateClientContact :one
INSERT INTO contacts (name, officephone, mobilephone, emailaddress, baseclient) VALUES ($1, $2, $3, $4, $5) 
RETURNING contactid;

-- name: DeleteClientContact :exec
DELETE FROM contacts WHERE contactid = $1;

-- name: ListClientContact :many
SELECT * FROM contacts WHERE baseclient = $1;

-- name: CreatePartnerContact :one
INSERT INTO contacts (name, officephone, mobilephone, emailaddress, basepartners) VALUES ($1, $2, $3, $4, $5)
RETURNING contactid;

-- name: DeletePartnerContact :exec
DELETE FROM contacts WHERE contactid = $1;

-- name: ListPartnerContact :many
SELECT * FROM contacts WHERE basepartners = $1;

-- name: CreateQualitymonitorContact :one
INSERT INTO contacts (name, officephone, mobilephone, emailaddress, baseqm) VALUES ($1, $2, $3, $4, $5)
RETURNING contactid;

-- name: DeleteQualitymonitorContact :exec
DELETE FROM contacts WHERE contactid = $1;

-- name: ListQualitymonitorContact :many
SELECT * FROM contacts WHERE baseqm = $1;


