package entities

import (
	"database/sql"
	"fmt"
	"time"
)

const SessionsTable = `
create table if not exists sessions (
	id integer primary key autoincrement,
	expiry timestamp not null,
	token text default '',
	user_id integer not null,
	foreign key (user_id) references users(id)
);`

type Session struct {
	UserId int
	Expiry time.Time
	Token  string
}

func (s *Session) Add(db *sql.DB) error {
	query := `
	INSERT INTO sessions (user_id, expiry, token)
	VALUES (?, ?, ?)`
	_, err := db.Exec(query, s.UserId, s.Expiry, s.Token)
	if err != nil {
		return fmt.Errorf("error adding session: %w", err)
	}
	return nil
}

func (s *Session) GetByToken(db *sql.DB, token string) (*Session, error) {
	query := `
	SELECT user_id, expiry, token
	FROM sessions
	WHERE token = ?`
	var session Session
	err := db.QueryRow(query, token).Scan(
		&session.UserId,
		&session.Expiry,
		&session.Token,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error querying session: %w", err)
	}
	return &session, nil
}
