// presentation/http/server/spec.go
package server

import "github.com/leandroluk/gox/oas"

type Spec struct {
	Path        string
	Method      string
	OperationFn func(*oas.Operation)
}
