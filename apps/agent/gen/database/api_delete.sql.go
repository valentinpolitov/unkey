// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: api_delete.sql

package database

import (
	"context"
)

const deleteApi = `-- name: DeleteApi :exec
DELETE FROM ` + "`" + `apis` + "`" + ` WHERE id = ?
`

func (q *Queries) DeleteApi(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteApi, id)
	return err
}