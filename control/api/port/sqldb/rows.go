// port/sqldb/rows.go
package sqldb

// Rows represents multiple rows of results
type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close() error
	Err() error
}
