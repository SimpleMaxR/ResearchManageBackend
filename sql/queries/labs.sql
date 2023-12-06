-- name: HealthzDatabase :one
SELECT version();


-- name: ListLab :one
SELECT * FROM Laboratories WHERE Name = $1;

-- name: ListLabAll :many
SELECT * FROM Laboratories;

-- name: CreateLab :one
INSERT INTO Laboratories (Name, OfficeArea, Address, ResearchDirection) VALUES ($1, $2, $3, $4) 
RETURNING LabID;

-- name: UpdateLab :one
UPDATE Laboratories SET Name = $1, OfficeArea = $2, Address = $3, ResearchDirection = $4 WHERE LabID = $5
RETURNING *;

-- name: DeleteLab :one
DELETE FROM Laboratories WHERE LabID = $1
RETURNING *;
