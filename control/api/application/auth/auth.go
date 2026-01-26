// application/auth/auth.go
package auth

import (
	"src/application/auth/activate_account"
	"src/application/auth/check_email_available"
	"src/application/auth/login_using_credential"
	"src/application/auth/refresh_token"
	"src/application/auth/register_account"
	"src/application/auth/resend_activation_code"
	"src/application/auth/reset_password"
	"src/application/auth/send_reset_password"
	"src/application/auth/sso_provider_redirect"
)

func Provide() {
	activate_account.Provide()
	check_email_available.Provide()
	login_using_credential.Provide()
	refresh_token.Provide()
	register_account.Provide()
	resend_activation_code.Provide()
	reset_password.Provide()
	send_reset_password.Provide()
	sso_provider_redirect.Provide()
}
