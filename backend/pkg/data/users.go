package data

import (
	"time"
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
    created_at timestamp default current_timestamp,
);
`

type User struct {
	Id        int       `json:"-"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Salt      string    `json:"-"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}
