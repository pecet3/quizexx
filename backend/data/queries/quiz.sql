
-- name: AddGameContents :one
insert into game_contents (uuid, user_id, max_rounds, category, gen_content, language, difficulty, content_json, created_at)
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
              RETURNING *;

-- name: GetGameContentByID :one
select * from game_contents where id = ?;


-- name: AddGameContentRound :one
INSERT INTO game_content_rounds (round, question_content, correct_answer_index, game_content_id)
              VALUES (?, ?, ?, ?)
              RETURNING *;


-- name: GetGameContentRoundByID :one
select * from game_content_rounds where id = ?;

-- name: AddGameRoundAnswer :one
INSERT INTO game_content_answers (is_correct, content, game_content_round_id)
              VALUES (?, ?, ?)
              RETURNING *;


-- name: AddGame :one
INSERT INTO games (room_uuid, room_name, game_content_id)
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateGame :one
UPDATE games
SET room_uuid = ?, room_name = ?, game_content_id = ?
WHERE id = ?
RETURNING *;

-- name: GetGame :one
SELECT * FROM games
WHERE id = ?;

-- name: GetGameByRoomUUID :one
SELECT * FROM games
WHERE room_uuid = ?;

-- name: AddGameWinner :one
INSERT INTO game_winner (points, game_id, user_id)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetGameWinner :one
SELECT * FROM game_winner
WHERE id = ?;

-- name: GetGameWinnersByGameID :many
SELECT * FROM game_winner
WHERE game_id = ?;

-- name: AddGameRoundAction :one
INSERT INTO game_round_action (answer_id, points, game_id, user_id)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: UpdateGameRoundAction :one
UPDATE game_round_action
SET answer_id = ?, points = ?, game_id = ?, user_id = ?
WHERE id = ?
RETURNING *;

-- name: GetGameRoundAction :one
SELECT * FROM game_round_action
WHERE id = ?;

-- name: GetGameRoundActionsByUserID :many
SELECT * FROM game_round_action
WHERE user_id = ?;

-- name: AddGameUser :one
INSERT INTO game_users (user_id, level, exp, games_wins, round_wins)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateGameUser :one
UPDATE game_users
SET level = ?, exp = ?, games_wins = ?, round_wins = ?
WHERE id = ?
RETURNING *;

-- name: GetGameUser :one
SELECT * FROM game_users
WHERE id = ?;

-- name: GetGameUserByUserID :one
SELECT * FROM game_users
WHERE user_id = ?;