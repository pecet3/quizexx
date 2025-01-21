--name: InsertGameContent :one
INSERT INTO game_contents (uuid, max_rounds, category, gen_content, language, difficulty, content_json, created_at)
              VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
              RETURNING *;

--name 