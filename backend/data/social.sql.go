// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: social.sql

package data

import (
	"context"
)

const addFunFact = `-- name: AddFunFact :one
INSERT INTO fun_facts (content, topic)
VALUES (?, ?)
RETURNING id, topic, content, created_at
`

type AddFunFactParams struct {
	Content string `json:"content"`
	Topic   string `json:"topic"`
}

func (q *Queries) AddFunFact(ctx context.Context, arg AddFunFactParams) (FunFact, error) {
	row := q.db.QueryRowContext(ctx, addFunFact, arg.Content, arg.Topic)
	var i FunFact
	err := row.Scan(
		&i.ID,
		&i.Topic,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const getCurrentFunFact = `-- name: GetCurrentFunFact :one
SELECT id, topic, content, created_at FROM fun_facts
ORDER BY created_at DESC
LIMIT 1
`

func (q *Queries) GetCurrentFunFact(ctx context.Context) (FunFact, error) {
	row := q.db.QueryRowContext(ctx, getCurrentFunFact)
	var i FunFact
	err := row.Scan(
		&i.ID,
		&i.Topic,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const getUsersSortedByExp = `-- name: GetUsersSortedByExp :many
SELECT u.id, u.uuid, u.name, u.email, u.salt, u.image_url, u.is_draft, u.created_at
FROM users u
JOIN game_users gu ON u.id = gu.user_id
ORDER BY gu.exp  DESC
LIMIT ? OFFSET ?
`

type GetUsersSortedByExpParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) GetUsersSortedByExp(ctx context.Context, arg GetUsersSortedByExpParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsersSortedByExp, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.Name,
			&i.Email,
			&i.Salt,
			&i.ImageUrl,
			&i.IsDraft,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
