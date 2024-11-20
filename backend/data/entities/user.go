package entities

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

const UsersTable = `
create table if not exists users (
    id integer primary key autoincrement,
    uuid text not null,
    name text default '',
    email text not null,
    password text default '',
    salt text not null,
    image_url text default '',
    created_at timestamp default current_timestamp
);
`

type User struct {
	ID        int       `json:"-"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Salt      string    `json:"-"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) Add(db *sql.DB) (int, error) {
	u.UUID = uuid.NewString()
	query := `INSERT INTO users (uuid, name, email, password, salt, image_url, created_at) 
              VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`

	result, err := db.Exec(query, u.UUID, u.Name, u.Email, u.Password, u.Salt, u.ImageUrl)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (u *User) Update(db *sql.DB) error {
	query := `UPDATE users SET name = ?, email = ?, password = ?, salt = ?, image_url = ? WHERE id = ?`
	_, err := db.Exec(query, u.Name, u.Email, u.Password, u.Salt, u.ImageUrl, u.ID)
	return err
}

func DeleteById(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no user found with the given ID")
	}

	return nil
}

func (u User) GetById(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, uuid, name, email, password, salt, image_url, created_at FROM users WHERE id = ?`
	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.Salt, &user.ImageUrl, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (u User) GetByEmail(db *sql.DB, email string) (*User, error) {
	query := `SELECT id, uuid, name, email, password, salt, image_url, created_at FROM users WHERE email = ?`
	row := db.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.Salt, &user.ImageUrl, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
