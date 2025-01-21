create table if not exists users (
    id integer primary key autoincrement,
    uuid text not null,
    name text not null default '',
    email text unique,
    salt text not null,
    image_url text default "/api/img/avatar.png",
	is_draft bool default true,
    created_at timestamp default current_timestamp
);

create table if not exists sessions (
	id integer primary key autoincrement,
	user_id integer not null,
	email text not null,
	expiry timestamp not null,
	token text default '',
	activate_code text default '',
	refresh_token text not null,
	user_ip text default '',
	type text not null,
	post_suspend_expiry timestamp,
	is_expired bool default false,
	foreign key (user_id) references users(id)
);