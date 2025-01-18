package entities

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type GameContent struct {
	ID          int
	UUID        string
	MaxRounds   int
	Category    string
	GenContent  string
	Language    string
	Difficulty  string
	ContentJSON string
	Round       GameContentRound
	Answer      GameContentAnswer
}

const GameContentTable = `
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
);`

func (gc *GameContent) Add(db *sql.DB) (int, error) {
	gc.UUID = uuid.NewString()
	query := `INSERT INTO game_contents (uuid, max_rounds, category, gen_content, language, difficulty, content_json, created_at)
              VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`

	result, err := db.Exec(query, gc.UUID, gc.MaxRounds, gc.Category, gc.GenContent, gc.Language, gc.Difficulty, gc.ContentJSON)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (gc *GameContent) Update(db *sql.DB) error {
	query := `UPDATE game_contents SET max_rounds = ?, category = ?, language = ?, content_json = ? WHERE id = ?`
	_, err := db.Exec(query, gc.MaxRounds, gc.Category, gc.Language, gc.ContentJSON, gc.ID)
	return err
}

func DeleteGameContentById(db *sql.DB, id int) error {
	query := `DELETE FROM game_contents WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no game content found with the given ID")
	}

	return nil
}

func (gc GameContent) GetById(db *sql.DB, id int) (*GameContent, error) {
	query := `SELECT id, uuid, max_rounds, category, language, content_json FROM game_contents WHERE id = ?`
	row := db.QueryRow(query, id)

	var content GameContent
	err := row.Scan(&content.ID, &content.UUID, &content.MaxRounds, &content.Category, &content.Language, &content.ContentJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("game content not found")
		}
		return nil, err
	}

	return &content, nil
}
