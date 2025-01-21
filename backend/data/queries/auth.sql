-- name: GetAllUsers :many
SELECT * FROM users;

-- name: InsertUser :one
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

-- name: DeleteUserByID: exec
DELETE FROM users WHERE id = ?;