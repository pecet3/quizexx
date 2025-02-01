-- name: AddFunFact :one
INSERT INTO fun_facts (content, topic)
VALUES (?, ?)
RETURNING *;

-- name: GetCurrentFunFact :one
SELECT * FROM fun_facts
ORDER BY created_at DESC
LIMIT 1;

-- name: GetUsersSortedByExp :many
SELECT u.*
FROM users u
JOIN game_users gu ON u.id = gu.user_id
ORDER BY gu.exp  DESC
LIMIT ? OFFSET ?;
