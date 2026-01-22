// application/auth/auth.go
package auth

import (
	"src/application/auth/activate_account"
	"src/application/auth/check_email_available"
	"src/application/auth/register_account"

	"github.com/leandroluk/gox/di"
)

func Provide() {
	di.SingletonAs[*activate_account.Handler](activate_account.New)
	di.SingletonAs[*check_email_available.Handler](check_email_available.New)
	di.SingletonAs[*register_account.Handler](register_account.New)
}
