-- name: ListSubtopicByProject :many
SELECT * from subtopics WHERE subtopics.projectid = $1;

-- name: ListSubtopicByLeader :many
SELECT * from subtopics WHERE subtopics.leaderid = $1;

-- name: CreateSubtopic :one
INSERT INTO subtopics (projectid, leaderid, name, enddate, fund, tech)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UpdateSubtopic :one
UPDATE subtopics SET projectid = $1, leaderid = $2, enddate = $3, fund = $4, tech = $5, name = $6 WHERE subtopicid = $7 RETURNING *;

-- name: DeleteSubtopic :exec
DELETE FROM subtopics WHERE subtopicid = $1;


