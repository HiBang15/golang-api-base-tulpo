// Code generated by sqlc. DO NOT EDIT.
// source: activity.sql

package database

import (
	"context"
	"database/sql"
)

const createActivity = `-- name: CreateActivity :one
INSERT INTO activities (
    url, method, url_regex
) VALUES (
    $1, $2, $3
)
RETURNING id, url, method, url_regex, created_at, deleted_at, updated_at
`

type CreateActivityParams struct {
	Url      string         `json:"url"`
	Method   sql.NullString `json:"method"`
	UrlRegex string         `json:"url_regex"`
}

func (q *Queries) CreateActivity(ctx context.Context, arg CreateActivityParams) (Activity, error) {
	row := q.db.QueryRowContext(ctx, createActivity, arg.Url, arg.Method, arg.UrlRegex)
	var i Activity
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Method,
		&i.UrlRegex,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteActivity = `-- name: DeleteActivity :exec
UPDATE activities
SET deleted_at = now()
WHERE id = $1
`

func (q *Queries) DeleteActivity(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteActivity, id)
	return err
}

const getActivityByID = `-- name: GetActivityByID :one
SELECT id, url, method, url_regex, created_at, deleted_at, updated_at FROM activities WHERE id = 1 AND deleted_at is null
`

func (q *Queries) GetActivityByID(ctx context.Context) (Activity, error) {
	row := q.db.QueryRowContext(ctx, getActivityByID)
	var i Activity
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Method,
		&i.UrlRegex,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getListActivity = `-- name: GetListActivity :many
SELECT id, url, method, url_regex, created_at, deleted_at, updated_at FROM activities WHERE deleted_at is null
`

func (q *Queries) GetListActivity(ctx context.Context) ([]Activity, error) {
	rows, err := q.db.QueryContext(ctx, getListActivity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Activity
	for rows.Next() {
		var i Activity
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.Method,
			&i.UrlRegex,
			&i.CreatedAt,
			&i.DeletedAt,
			&i.UpdatedAt,
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

const updateActivity = `-- name: UpdateActivity :one
UPDATE activities
SET url = $2, method = $3, url_regex = $4
WHERE id = $1 and deleted_at is null
RETURNING id, url, method, url_regex, created_at, deleted_at, updated_at
`

type UpdateActivityParams struct {
	ID       int32          `json:"id"`
	Url      string         `json:"url"`
	Method   sql.NullString `json:"method"`
	UrlRegex string         `json:"url_regex"`
}

func (q *Queries) UpdateActivity(ctx context.Context, arg UpdateActivityParams) (Activity, error) {
	row := q.db.QueryRowContext(ctx, updateActivity,
		arg.ID,
		arg.Url,
		arg.Method,
		arg.UrlRegex,
	)
	var i Activity
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Method,
		&i.UrlRegex,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.UpdatedAt,
	)
	return i, err
}
