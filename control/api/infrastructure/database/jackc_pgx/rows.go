// infrastructure/database/jackc_pgx/rows.go
package jackc_pgx

import (
	sqldb "src/port/database"

	"github.com/jackc/pgx/v5"
)

// Rows
type Rows struct {
	rows pgx.Rows
}

var _ sqldb.Rows = (*Rows)(nil)

func (r *Rows) Next() bool {
	return r.rows.Next()
}

func (r *Rows) Scan(dest ...any) error {
	return r.rows.Scan(dest...)
}

func (r *Rows) Close() error {
	r.rows.Close()
	return nil
}

func (r *Rows) Err() error {
	return r.rows.Err()
}
