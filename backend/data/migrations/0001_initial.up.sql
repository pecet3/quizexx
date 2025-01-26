CREATE TABLE IF NOT EXISTS game_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER UNIQUE NOT NULL,
    level INTEGER NOT NULL DEFAULT 1,
    exp INTEGER NOT NULL DEFAULT 0,
    games_wins INTEGER NOT NULL DEFAULT 0,
    round_wins INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);


create table if not exists users (
    id integer primary key autoincrement,
    uuid text not null,
    name text not null not null,
    email text unique,
    salt text not null,
    image_url text not null,
	is_draft bool not null,
    created_at timestamp default current_timestamp
);

create table if not exists sessions (
	id integer primary key autoincrement,
	user_id integer not null,
	email text not null,
	expiry timestamp not null,
	token text not null,
	activate_code text not null,
	refresh_token text not null,
	user_ip text not null,
	type text not null,
	post_suspend_expiry timestamp,
	is_expired bool default false,
	foreign key (user_id) references users(id)
);

CREATE TABLE IF NOT EXISTS game_contents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT NOT NULL UNIQUE,
    max_rounds INTEGER NOT NULL,
    category TEXT NOT NULL,
    gen_content TEXT NOT NULL,
    language TEXT NOT NULL,
    difficulty TEXT NOT NULL,
    content_json TEXT NOT NULL,
	user_id integer not null,
    created_at timestamp default current_timestamp,
	foreign key (user_id) references users(id)
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