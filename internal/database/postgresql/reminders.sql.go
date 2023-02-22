// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: reminders.sql

package postgresql

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getAllReminders = `-- name: GetAllReminders :many
SELECT id, title, description, status, created FROM reminders
ORDER BY created ASC
`

func (q *Queries) GetAllReminders(ctx context.Context) ([]Reminder, error) {
	rows, err := q.db.QueryContext(ctx, getAllReminders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reminder
	for rows.Next() {
		var i Reminder
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Created,
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

const getRemindersWithId = `-- name: GetRemindersWithId :many
SELECT id, title, description, status, created FROM reminders 
WHERE id = ANY($1::uuid[])
ORDER BY created ASC
`

func (q *Queries) GetRemindersWithId(ctx context.Context, ids []uuid.UUID) ([]Reminder, error) {
	rows, err := q.db.QueryContext(ctx, getRemindersWithId, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reminder
	for rows.Next() {
		var i Reminder
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Created,
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
