// presentation/http/route/auth/auth.go
package auth

import (
	"src/presentation/http/route/auth/activate_account"
	"src/presentation/http/route/auth/check_email_available"
	"src/presentation/http/route/auth/login_using_code"
	"src/presentation/http/route/auth/login_using_credential"
	"src/presentation/http/route/auth/refresh_token"
	"src/presentation/http/route/auth/register_account"
	"src/presentation/http/route/auth/resend_activation_code"
	"src/presentation/http/route/auth/reset_password"
	"src/presentation/http/route/auth/send_reset_password"
	"src/presentation/http/route/auth/sso_provider_callback"
	"src/presentation/http/route/auth/sso_provider_redirect"
	"src/presentation/http/server"
)

var Group = server.
	Group("/auth", func(gr *server.Grouper) {
		gr.Patch("/activate", activate_account.Route)
		gr.Post("/activate", resend_activation_code.Route)
		gr.Get("/check/email/:email", check_email_available.Route)
		gr.Post("/login/code", login_using_code.Route)
		gr.Post("/login/credential", login_using_credential.Route)
		gr.Patch("/recover", reset_password.Route)
		gr.Post("/recover", send_reset_password.Route)
		gr.Post("/refresh", refresh_token.Route)
		gr.Post("/register", register_account.Route)
		gr.Get("/:provider/callback", sso_provider_callback.Route)
		gr.Get("/:provider", sso_provider_redirect.Route)
	})
