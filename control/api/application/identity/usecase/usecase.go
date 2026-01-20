package usecase

import (
	"src/application/identity/usecase/check_email_available"

	"github.com/leandroluk/gox/di"
)

func Provide() {
	di.SingletonAs[*check_email_available.Handler](check_email_available.New)
}
