-- name: GetAllUsers :many
SELECT * FROM users;

-- name: AddUser :one
INSERT INTO users (uuid, name, email, salt, image_url, is_draft, created_at)
              VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
              RETURNING *;

-- name: GetUserByID :one
SELECT * from users WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * from users WHERE email = ? LIMIT 1;

-- name: UpdateUserName :one
UPDATE users SET name = ? WHERE id = ?
    RETURNING *;

-- name: UpdateUserIsDraft :one
UPDATE users SET is_draft = ? WHERE id = ?
    RETURNING *;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = ?;

-- name: AddSession :one
INSERT INTO sessions (
    user_id, email, expiry, token, activate_code, user_ip, type, post_suspend_expiry, is_expired
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetSessionByToken :one
SELECT *
FROM sessions
WHERE token = ?;

-- name: UpdatePostSuspendExpiry :exec
UPDATE sessions
SET post_suspend_expiry = ?
WHERE token = ?;

-- name: UpdateIsExpired :exec
UPDATE sessions
SET is_expired = ?
WHERE token = ?;
