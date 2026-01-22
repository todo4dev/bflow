// infrastructure/database/jackc_pgx/result.go
package jackc_pgx

import (
	sqldb "src/port/database"

	"github.com/jackc/pgx/v5/pgconn"
)

// Result
type Result struct {
	tag pgconn.CommandTag
}

var _ sqldb.Result = (*Result)(nil)

func (r *Result) LastInsertId() (int64, error) {
	return 0, nil // Postgres doesn't support LastInsertId naturally in the same way MySQL does. Usually requires RETURNING id.
}

func (r *Result) RowsAffected() (int64, error) {
	return r.tag.RowsAffected(), nil
}
