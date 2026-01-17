// port/sqldb/row.go
package sqldb

// Row represents a single row of results
type Row interface {
	Scan(dest ...any) error
}
