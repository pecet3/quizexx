CREATE TABLE IF NOT EXISTS game_contents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT NOT NULL UNIQUE,
    max_rounds INTEGER NOT NULL,
    category TEXT NOT NULL,
    gen_content TEXT NOT NULL,
    language TEXT NOT NULL,
    difficulty TEXT NOT NULL,
    content_json TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS game_content_rounds (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    round INTEGER NOT NULL,
    question_content TEXT NOT NULL,
    correct_answer_index INTEGER NOT NULL,
    game_content_id INTEGER NOT NULL,
    FOREIGN KEY (game_content_id) REFERENCES game_contents(id)
);

CREATE TABLE IF NOT EXISTS game_content_answers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    is_correct BOOLEAN NOT NULL,
    content TEXT NOT NULL,
    game_content_round_id INTEGER NOT NULL,
    FOREIGN KEY (game_content_round_id) REFERENCES game_content_rounds(id)
);