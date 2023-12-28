-- name: Login :one
SELECT * FROM Users WHERE Username = $1 AND Password = $2 LIMIT 1;