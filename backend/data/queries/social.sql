-- name: AddFunFact :one
INSERT INTO fun_facts (content, topic)
VALUES (?, ?)
RETURNING *;

-- name: GetCurrentFunFact :one
SELECT * FROM fun_facts
ORDER BY created_at DESC
LIMIT 1;