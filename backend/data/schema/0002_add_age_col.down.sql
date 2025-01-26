CREATE TEMPORARY TABLE users_backup AS SELECT id, uuid, name, email salt, image_url, is_draft FROM users;
DROP TABLE users;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE

);
INSERT INTO users (id, name, email) SELECT id, name, email FROM users_backup;
DROP TABLE users_backup;

create table if not exists users (
    id integer primary key autoincrement,
    uuid text not null,
    name text not null default '',
    email text unique,
    salt text not null,
    image_url text default "/api/img/avatar.png",
	is_draft bool default true,
    created_at timestamp default current_timestamp
)