// application/auth/auth.go
package auth

import (
	"src/application/auth/activate_account"
	"src/application/auth/check_email_available"
	"src/application/auth/login_using_credential"
	"src/application/auth/refresh_authorization_token"
	"src/application/auth/register_account"
	"src/application/auth/resend_activation_code"
	"src/application/auth/reset_password"
	"src/application/auth/send_reset_password"

	"github.com/leandroluk/gox/di"
)

func Provide() {
	di.SingletonAs[*activate_account.Handler](activate_account.New)
	di.SingletonAs[*check_email_available.Handler](check_email_available.New)
	di.SingletonAs[*login_using_credential.Handler](login_using_credential.New)
	di.SingletonAs[*refresh_authorization_token.Handler](refresh_authorization_token.New)
	di.SingletonAs[*register_account.Handler](register_account.New)
	di.SingletonAs[*resend_activation_code.Handler](resend_activation_code.New)
	di.SingletonAs[*reset_password.Handler](reset_password.New)
	di.SingletonAs[*send_reset_password.Handler](send_reset_password.New)
}
