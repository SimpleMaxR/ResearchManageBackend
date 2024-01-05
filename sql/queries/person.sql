-- Researcher queries

-- name: ListResearcherAll :many
SELECT * FROM Researchers;

-- name: ListResearcher :one
SELECT * FROM Researchers WHERE ResearcherID = $1;

-- name: CreateResearcher :one
INSERT INTO Researchers (lab_id, researcher_number, Name, Gender, Title, Age, emailaddress, Leader, startdate, term, researchDirection) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING ResearcherID;

-- name: UpdateResearcher :one
UPDATE Researchers SET lab_id = $1, Name = $2, Gender = $3, Title = $4, Age = $5, ResearchDirection = $6, Leader = $7 WHERE ResearcherID = $8
RETURNING *;

-- name: DeleteResearcher :one
DELETE FROM Researchers WHERE ResearcherID = $1
RETURNING *;

-- name: ListResearcherByLab :many
SELECT * FROM Researchers WHERE lab_id = $1;

-- name: ListResearcherByID :one
SELECT * FROM Researchers WHERE ResearcherID = $1;


-- -- Director queries

-- -- name: ListDirectorAll :many
-- SELECT * FROM Directors;

-- -- name: ListDirector :one
-- SELECT * FROM Directors WHERE DirectorID = $1;

-- -- name: CreateDirector :one
-- INSERT INTO Directors (lab_id, StartDate, Term) VALUES ($1, $2, $3)
-- RETURNING DirectorID;

-- -- name: setResearcherDirector :one
-- UPDATE Researchers SET Leader = true WHERE ResearcherID = $1
-- RETURNING *;

-- -- name: unsetResearcherDirector :one
-- UPDATE Researchers SET Leader = false WHERE ResearcherID = $1
-- RETURNING *;

-- -- name: UpdateDirector :one
-- UPDATE Directors SET lab_id = $1, StartDate = $2, Term = $3 WHERE DirectorID = $4
-- RETURNING *;

