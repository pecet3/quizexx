package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pecet3/quizex/data/entities"
)

const PREFIX = "/v1"

type Data struct {
	Db          *sql.DB
	User        entities.User
	Session     entities.Session
	GameContent entities.GameContent
}

func NewData() *Data {
	return &Data{
		Db:      NewSQLite(),
		User:    entities.User{},
		Session: entities.Session{},
		GameContent: entities.GameContent{
			Round:  entities.GameContentRound{},
			Answer: entities.GameContentAnswer{},
		},
	}
}
