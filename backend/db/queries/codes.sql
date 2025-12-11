-- name: CreateCode :one 
INSERT INTO codes (
  code,
  user_id,
  expires_at
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetCode :one
SELECT * FROM codes WHERE user_id = $1 AND code = $2 AND used = FALSE AND expires_at > NOW() LIMIT 1;

-- name: MarkCodeUsed :exec
UPDATE codes SET used = TRUE where id = $1 AND used = FALSE;

-- name: DeleteCode :exec
DELETE FROM codes WHERE id = $1;

-- name: DeleteExpiredCodes :exec
DELETE FROM codes WHERE expires_at < NOW();
