
-- name: AddGameContents :one
insert into game_contents (uuid, max_rounds, category, gen_content, language, difficulty, content_json, created_at)
              VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
              RETURNING *;

-- name: GetGameContentByID :one
select * from game_contents where id = ?;


-- name: AddGameContentRound :one
INSERT INTO game_content_rounds (round, question_content, correct_answer_index, game_content_id)
              VALUES (?, ?, ?, ?)
              RETURNING *;


-- name: GetGameContentRoundByID :one
select * from game_content_rounds where id = ?;

-- name: AddGameRound :one
INSERT INTO game_content_answers (is_correct, content, game_content_round_id)
              VALUES (?, ?, ?)
              RETURNING *;

              