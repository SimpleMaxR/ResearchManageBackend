-- achievements

-- name: CreateAchievement :one
INSERT INTO achievements (name, obtaineddate, contributorid, baseproject, basesubtopic, type) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING achievementid;

-- name: DeleteAchievement :exec
DELETE FROM achievements WHERE achievementid = $1;

-- name: ListAchievementByProject :many
SELECT * FROM achievements WHERE baseproject = $1;

-- name: ListAchievementBySubtopic :many
SELECT * FROM achievements WHERE basesubtopic = $1;