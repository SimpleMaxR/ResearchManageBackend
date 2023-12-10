-- name: AdminLogin :one
SELECT * FROM Users WHERE Role = 'admin' AND Username = $1 AND Password = $2 LIMIT 1;