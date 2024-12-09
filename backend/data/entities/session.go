package entities

import (
	"database/sql"
	"fmt"
	"time"
)

const SessionsTable = `
create table if not exists sessions (
	id integer primary key autoincrement,
	user_id integer not null,
	email text not null,
	expiry timestamp not null,
	token text default '',
	activate_code text default '',
	user_ip text default '',
	type text not null,
	post_suspend_expiry timestamp,
	is_expired bool default false,
	foreign key (user_id) references users(id)
);`

type Session struct {
	UserId            int
	Email             string
	Expiry            time.Time
	Token             string
	ActivateCode      string
	UserIp            string
	Type              string
	IsExpired         bool
	PostSuspendExpiry time.Time
}

func (s Session) Add(db *sql.DB, session *Session) error {
	query := `insert into sessions (user_id,
	 email, expiry, token, activate_code, user_ip, type, post_suspend_expiry, is_expired)
	  values (?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(query,
		session.UserId,
		session.Email,
		session.Expiry,
		session.Token,
		session.ActivateCode,
		session.UserIp,
		session.Type,
		session.PostSuspendExpiry,
		session.IsExpired,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s Session) GetByToken(db *sql.DB, token string) (*Session, error) {
	query := "SELECT user_id, email, expiry, token, activate_code, user_ip, type, post_suspend_expiry, is_expired FROM sessions WHERE token = ?"
	var session Session
	err := db.QueryRow(query, token).Scan(
		&session.UserId,
		&session.Email,
		&session.Expiry,
		&session.Token,
		&session.ActivateCode,
		&session.UserIp,
		&session.Type,
		&session.PostSuspendExpiry,
		&session.IsExpired,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Session not found, but it's not an error
		}
		return nil, fmt.Errorf("error querying session: %w", err)
	}
	return &session, nil
}

func (s Session) UpdatePostSuspendExpiry(db *sql.DB, token string, newExpiry time.Time) error {
	query := `UPDATE sessions SET post_suspend_expiry = ? WHERE token = ?`
	_, err := db.Exec(query, newExpiry, token)
	if err != nil {
		return fmt.Errorf("error updating post_suspend_expiry: %w", err)
	}
	return nil
}

func (s Session) UpdateIsExpired(db *sql.DB, token string, isExpired bool) error {
	query := `UPDATE sessions SET is_expired = ? WHERE token = ?`
	_, err := db.Exec(query, isExpired, token)
	if err != nil {
		return fmt.Errorf("error updating is_expired: %w", err)
	}
	return nil
}
