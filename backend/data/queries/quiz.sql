
-- name: AddGameContents :one
insert into game_contents (uuid, user_id, max_rounds, category, gen_content, language, difficulty, content_json, created_at)
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
              RETURNING *;

-- name: GetGameContentByID :one
select * from game_contents where id = ?;

-- name: GetGameContentByUUID :one
select * from game_contents where uuid = ?;


-- name: AddGameContentRound :one
INSERT INTO game_content_rounds (round, question_content, correct_answer_index, game_content_id)
              VALUES (?, ?, ?, ?)
              RETURNING *;

-- name: GetGameContentRoundsByGameUUID :one
select * from game_content_rounds where game_content_id = ?;

-- name: GetGameContentRoundByID :one
select * from game_content_rounds where id = ?;

-- name: GetGameContentRoundsByGameContentID :many
SELECT * FROM game_content_rounds where game_content_id = ?;

-- name: AddGameRoundAnswer :one
INSERT INTO game_content_answers (is_correct, content, round_number, game_content_round_id, index_in_arr)
              VALUES (?, ?, ?, ?, ?)
              RETURNING *;

-- name: GetGameContentAnswerByRoundIDAndIndex :one
SELECT * FROM game_content_answers where game_content_round_id = ? and index_in_arr = ?;


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
INSERT INTO game_winners (points, game_id, user_id)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetGameWinner :one
SELECT * FROM game_winners
WHERE id = ?;

-- name: GetGameWinnersByGameID :many
SELECT * FROM game_winners
WHERE game_id = ?;

-- name: AddGameRoundAction :one
INSERT INTO game_round_actions (answer_id, game_content_round_id, answer_index, is_good_answer, points, game_id, user_id)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdatePointsRoundAction :one
UPDATE game_round_actions
SET points = ?
WHERE game_id = ? and
game_content_round_id = ?
RETURNING *;

-- name: UpdateGameRoundAction :one
UPDATE game_round_actions
SET answer_id = ?, points = ?, game_id = ?, user_id = ?
WHERE id = ?
RETURNING *;

-- name: GetGameRoundAction :one
SELECT * FROM game_round_actions
WHERE id = ?;

-- name: GetGameRoundActionsByUserIDRoundIDGameID :one
SELECT * FROM game_round_actions
WHERE user_id = ? and game_content_round_id = ? and game_id = ?;

-- name: AddGameUser :one
INSERT INTO game_users (user_id, level, exp, games_wins, round_wins, progress)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateGameUser :one
UPDATE game_users
SET level = ?, exp = ?, games_wins = ?, round_wins = ?, progress = ?
WHERE id = ?
RETURNING *;

-- name: GetGameUser :one
SELECT * FROM game_users
WHERE id = ?;


-- name: GetGameUserByUserID :one
SELECT * FROM game_users
WHERE user_id = ?;