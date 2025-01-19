package entities

import (
	"database/sql"
	"errors"
)

type GameContentAnswer struct {
	ID                 int
	IsCorrect          bool
	Content            string
	GameContentRoundID int
}

const GameContentAnswerTable = `
CREATE TABLE IF NOT EXISTS game_content_answers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    is_correct BOOLEAN NOT NULL,
    content TEXT NOT NULL,
    game_content_round_id INTEGER NOT NULL,
    FOREIGN KEY (game_content_round_id) REFERENCES game_content_rounds(id)
);`

func (gca *GameContentAnswer) Add(db *sql.DB) (int, error) {
	query := `INSERT INTO game_content_answers (is_correct, content, game_content_round_id)
              VALUES (?, ?, ?, ?)`

	result, err := db.Exec(query, gca.IsCorrect, gca.Content, gca.GameContentRoundID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (gca *GameContentAnswer) Update(db *sql.DB) error {
	query := `UPDATE game_content_answers SET is_correct = ?, content = ? WHERE id = ?`
	_, err := db.Exec(query, gca.IsCorrect, gca.Content, gca.ID)
	return err
}

func (gca *GameContent) DeleteGameContentAnswerById(db *sql.DB, id int) error {
	query := `DELETE FROM game_content_answers WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no game content answer found with the given ID")
	}

	return nil
}
