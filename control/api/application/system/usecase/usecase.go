package usecase

import (
	"src/application/system/usecase/healthcheck"

	"github.com/leandroluk/gox/di"
)

func Provide() {
	di.SingletonAs[*healthcheck.Handler](healthcheck.New)
}
