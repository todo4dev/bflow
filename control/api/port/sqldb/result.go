// port/sqldb/result.go
package sqldb

// Result represents the result of a write operation
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
