-- name: CountLab :one
select count(*) from laboratories;

-- name: CountResearcher :one
select count(*) from researchers;

-- name: CountProject :one
select count(*) from projects;