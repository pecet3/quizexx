package entities

import (
	"database/sql"
	"errors"
)

type GameContentRound struct {
	ID                 int
	Round              int
	QuestionContent    string
	CorrectAnswerIndex int
	GameContentID      int
}

const GameContentRoundTable = `
CREATE TABLE IF NOT EXISTS game_content_rounds (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    round INTEGER NOT NULL,
    question_content TEXT NOT NULL,
    correct_answer_index INTEGER NOT NULL,
    game_content_id INTEGER NOT NULL,
    FOREIGN KEY (game_content_id) REFERENCES game_contents(id)
);`

func (gcr *GameContentRound) Add(db *sql.DB) (int, error) {
	query := `INSERT INTO game_content_rounds (round, question_content, correct_answer_index, game_content_id)
              VALUES (?, ?, ?, ?)`

	result, err := db.Exec(query, gcr.Round, gcr.QuestionContent, gcr.CorrectAnswerIndex, gcr.GameContentID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (gcr *GameContentRound) Update(db *sql.DB) error {
	query := `UPDATE game_content_rounds SET round = ?, question_content = ?, correct_answer_index = ? WHERE id = ?`
	_, err := db.Exec(query, gcr.Round, gcr.QuestionContent, gcr.CorrectAnswerIndex, gcr.ID)
	return err
}

func DeleteGameContentRoundById(db *sql.DB, id int) error {
	query := `DELETE FROM game_content_rounds WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no game content round found with the given ID")
	}

	return nil
}
