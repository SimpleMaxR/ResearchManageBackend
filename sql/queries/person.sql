-- Researcher queries

-- name: ListResearcherAll :many
SELECT * FROM Researchers;

-- name: ListResearcher :one
SELECT * FROM Researchers WHERE ResearcherID = $1;

-- name: CreateResearcher :one
INSERT INTO Researchers (lab_id, Name, Gender, Title, Age, ResearchDirection, Leader) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING ResearcherID;

-- name: UpdateResearcher :one
UPDATE Researchers SET lab_id = $1, Name = $2, Gender = $3, Title = $4, Age = $5, ResearchDirection = $6, Leader = $7 WHERE ResearcherID = $8
RETURNING *;

-- name: DeleteResearcher :one
DELETE FROM Researchers WHERE ResearcherID = $1
RETURNING *;

-- name: ListResearcherByLab :many
SELECT * FROM Researchers WHERE lab_id = $1;


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


-- Secretary queries

-- name: ListSecretaryAll :many
SELECT * FROM Secretaries;

-- name: CreateSecretaries :one
INSERT INTO Secretaries (Name, Gender, Age, MobilePhone, EmailAddress) VALUES ($1, $2, $3, $4, $5)
RETURNING SecretaryID;

-- name: CreateSecretaryServices :one
INSERT INTO SecretaryServices (SecretaryID, lab_id, EmploymentDate, Responsibilities) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateSecretaries :one
UPDATE Secretaries SET Name = $1, Gender = $2, Age = $3, MobilePhone = $4, EmailAddress = $5 WHERE SecretaryID = $6
RETURNING *;

-- name: UpdateSecretaryService :one
UPDATE SecretaryServices SET EmploymentDate = $1, Responsibilities = $2 WHERE SecretaryID = $3 AND lab_id = $4
RETURNING *;

-- name: DeleteSecretary :exec
DELETE FROM Secretaries WHERE SecretaryID = $1;
DELETE FROM SecretaryServices WHERE SecretaryID = $1;

-- name: DeleteSecretaryService :exec
DELETE FROM SecretaryServices WHERE SecretaryID = $1 AND lab_id = $2;

-- name: ListSecretaryServiceBySID :many
SELECT * FROM SecretaryServices WHERE SecretaryID = $1;

-- name: ListSecretaryServiceByLab :many
SELECT * FROM SecretaryServices WHERE lab_id = $1;
