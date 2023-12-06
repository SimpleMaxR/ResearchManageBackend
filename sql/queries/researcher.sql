-- name: ListResearcherAll :many
SELECT * FROM Researchers;

-- name: ListResearcher :one
SELECT * FROM Researchers WHERE ResearcherID = $1;

-- name: CreateResearcher :one
INSERT INTO Researchers (LabID, Name, Gender, Title, Age, ResearchDirection, Leader) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING ResearcherID;

-- name: UpdateResearcher :one
UPDATE Researchers SET LabID = $1, Name = $2, Gender = $3, Title = $4, Age = $5, ResearchDirection = $6, Leader = $7 WHERE ResearcherID = $8
RETURNING *;

-- name: DeleteResearcher :one
DELETE FROM Researchers WHERE ResearcherID = $1
RETURNING *;

-- name: ListResearcherByLab :many
SELECT * FROM Researchers WHERE LabID = $1;