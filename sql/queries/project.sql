-- name: CreateProject :one
INSERT INTO projects (projectleader, name, researchcontent, totalfunds, startdate, enddate, qualitymonitorsid, clientid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING projectid;

-- name: UpdateProject :one
UPDATE projects SET projectleader = $2, name = $3, researchcontent = $4, totalfunds = $5, startdate = $6, enddate = $7, qualitymonitorsid = $8, clientid = $9 WHERE projectid = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects WHERE projectid = $1;

-- name: ListProjectAll :many
SELECT * FROM projects;

-- name: GetProjectById :one
SELECT * FROM projects WHERE projectid = $1;

-- name: GetProjectByName :many
SELECT * FROM projects WHERE Name LIKE '%' || $1 || '%';


-- ProjectPartner

-- name: LinkProjectPartner :exec
INSERT INTO projectpartners (projectid, partnerid) VALUES ($1, $2);

-- name: UnlinkProjectPartner :exec
DELETE FROM projectpartners WHERE projectid = $1 AND partnerid = $2;

-- name: GetParterByProject :many
SELECT * FROM partners WHERE partnerid IN (SELECT partnerid FROM projectpartners WHERE projectid = $1) ORDER BY partnerid;



-- ProjectResearcher

-- name: LinkProjectResearcher :exec
INSERT INTO projectResearchers (projectid, researcherid, joindate, workload) VALUES ($1, $2, $3, $4);

-- name: UnlinkProjectResearcher :exec
DELETE FROM projectResearchers WHERE projectid = $1 AND researcherid = $2;

-- name: ListProjectResearcher :many
SELECT researcherid FROM projectResearchers WHERE projectid = $1;