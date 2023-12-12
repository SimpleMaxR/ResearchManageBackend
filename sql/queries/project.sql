-- name: CreateProject :one
INSERT INTO projects (peojectleader, name, researchcontent, totalfunds, startdate, enddate, qualitymonitorsid, clientid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING projectid;

-- name: UpdateProject :one
UPDATE projects SET peojectleader = $2, name = $3, researchcontent = $4, totalfunds = $5, startdate = $6, enddate = $7, qualitymonitorsid = $8, clientid = $9 WHERE projectid = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects WHERE projectid = $1;

-- name: ListProjectAll :many
SELECT * FROM projects;


-- ProjectPartner

-- name: LinkProjectPartner :exec
INSERT INTO projectpartners (projectid, partnerid) VALUES ($1, $2);

-- name: UnlinkProjectPartner :exec
DELETE FROM projectpartners WHERE projectid = $1 AND partnerid = $2;

-- name: ListProjectPartner :many
SELECT * FROM partners WHERE partnerid IN (SELECT partnerid FROM projectpartners WHERE projectid = $1);


-- ProjectResearcher

-- name: LinkProjectResearcher :exec
INSERT INTO projectResearchers (projectid, researcherid, joindate, workload, disposablefunds) VALUES ($1, $2, $3, $4, $5);

-- name: UnlinkProjectResearcher :exec
DELETE FROM projectResearchers WHERE projectid = $1 AND researcherid = $2;

-- name: ListProjectResearcher :many
SELECT * FROM researchers WHERE researcherid IN (SELECT researcherid FROM projectResearchers WHERE projectid = $1);


-- Subtopic

-- name: CreateSubtopic :one
INSERT INTO subtopics (projectid, leaderid, enddaterequirement, disposablefunds, technicalindicators) VALUES ($1, $2, $3, $4, $5)
RETURNING subtopicid;

-- name: UpdateSubtopic :one
UPDATE subtopics SET projectid = $2, leaderid = $3, enddaterequirement = $4, disposablefunds = $5, technicalindicators = $6 WHERE subtopicid = $1
RETURNING *;

-- name: DeleteSubtopic :exec
DELETE FROM subtopics WHERE subtopicid = $1;

-- name: ListSubtopic :many
SELECT * FROM subtopics WHERE projectid = $1;


-- achievements

-- name: CreateAchievement :one
INSERT INTO achievements (name, obtaineddate, contributorid, baseproject, basesubtopic, rank) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING achievementid;

-- name: DeleteAchievement :exec
DELETE FROM achievements WHERE achievementid = $1;

-- name: ListAchievement :many
SELECT * FROM achievements WHERE baseproject = $1;

-- name: ListAchievementBySubtopic :many
SELECT * FROM achievements WHERE basesubtopic = $1;