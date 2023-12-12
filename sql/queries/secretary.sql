-- name: CreateSecretary :one
INSERT INTO secretaries (name, gender, age, mobilephone, emailaddress) VALUES ($1, $2, $3, $4, $5) RETURNING secretaryid;

-- name: CreateSecretaryService :one
INSERT INTO secretaryservices (secretaryid, lab_id, employmentdate, responsibilities) VALUES ($1, $2, $3, $4) 
RETURNING secretaryid;

-- name: DeleteSecretary :exec
DELETE FROM secretaries WHERE secretaryid = $1;

-- name: DeleteSecretaryService :exec
DELETE FROM secretaryservices WHERE secretaryid = $1;

-- name: ListSecretaryAll :many
SELECT * FROM Secretaries;

-- name: ListSecretaryServiceBySID :many
SELECT * FROM SecretaryServices WHERE SecretaryID = $1;

-- name: ListSecretaryServiceByLab :many
SELECT * FROM SecretaryServices WHERE lab_id = $1;
