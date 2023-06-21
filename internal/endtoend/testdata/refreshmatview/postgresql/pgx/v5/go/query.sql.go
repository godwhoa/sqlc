// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"
)

const refresh = `-- name: Refresh :exec
REFRESH MATERIALIZED VIEW myview
`

func (q *Queries) Refresh(ctx context.Context) error {
	_, err := q.db.Exec(ctx, refresh)
	return err
}
