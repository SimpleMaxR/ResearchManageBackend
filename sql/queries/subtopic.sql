-- name: CreateSubtopic :one
INSERT INTO subtopics (projectid, leaderid, enddaterequirement, disposablefunds, technicalindicators) VALUES ($1, $2, $3, $4, $5) RETURNING subtopicid;

-- name: UpdateSubtopic :one
UPDATE subtopics SET projectid=$2, leaderid=$3, enddaterequirement=$4, disposablefunds=$5, technicalindicators=$6 WHERE subtopicid=$1 RETURNING *;

-- name: DeleteSubtopic :exec
DELETE FROM subtopics WHERE subtopicid=$1;

-- name: GetSubtopic :one
SELECT * FROM subtopics WHERE subtopicid=$1;

-- name: GetSubtopicByProject :many
SELECT * FROM subtopics WHERE projectid=$1;

-- name: GetSubtopicByLeader :many
SELECT * FROM subtopics WHERE leaderid=$1;