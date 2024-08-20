-- name: CreateUser :one
INSERT INTO users (user_name, user_password, user_email, user_phone, user_token)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE user_id = $1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: UpdateUser :one
UPDATE users
SET user_name = $2, user_password = $3, user_email = $4, user_phone = $5, user_token = $6
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;