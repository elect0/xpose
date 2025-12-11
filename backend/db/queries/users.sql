-- name: CreateUser :one
INSERT INTO users (
  email
) VALUES ($1) RETURNING *;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: UpdateUserById :one
UPDATE users SET username = $2, email = $3, profile_pic_url = $4, verified = $5 WHERE id = $1 RETURNING *;

-- name: VerifyUserById :one
UPDATE users SET verified = TRUE WHERE id = $1 RETURNING *;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;
