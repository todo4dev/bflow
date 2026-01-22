// application/system/system.go
package system

import (
	"src/application/system/healthcheck"

	"github.com/leandroluk/gox/di"
)

func Provide() {
	di.SingletonAs[*healthcheck.Handler](healthcheck.New)
}
