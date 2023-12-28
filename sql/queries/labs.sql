-- name: HealthzDatabase :one
SELECT version();


-- name: ListLab :one
SELECT * FROM laboratories WHERE Name = $1;

-- name: ListLabById :one
SELECT * FROM laboratories WHERE lab_id = $1;

-- name: ListLabAll :many
SELECT * FROM Laboratories;

-- name: CreateLab :one
INSERT INTO Laboratories (Name, office_area, Address, research_direction) VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: UpdateLab :one
UPDATE Laboratories SET Name = $1, office_area = $2, Address = $3, research_direction = $4 WHERE lab_id = $5
RETURNING *;

-- name: DeleteLab :one
DELETE FROM Laboratories WHERE lab_id = $1
RETURNING *;

-- name: ListDirectorByLab :one
SELECT * FROM Researchers WHERE lab_id = $1 AND Leader = true;

-- name: ListSecretaryByLab :one
SELECT * FROM SecretaryServices WHERE lab_id = $1;

-- name: ListOfficeByLab :one
SELECT * FROM Offices WHERE lab_id = $1;

