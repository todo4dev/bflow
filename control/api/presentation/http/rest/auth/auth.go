// presentation/http/rest/auth/auth.go
package auth

import (
	"src/presentation/http/rest/auth/activate_account"
	"src/presentation/http/rest/auth/check_email_available"
	"src/presentation/http/rest/auth/login_using_credential"
	"src/presentation/http/rest/auth/refresh_authorization_token"
	"src/presentation/http/rest/auth/register_account"
	"src/presentation/http/rest/auth/resend_activation_code"
	"src/presentation/http/rest/auth/reset_password"
	"src/presentation/http/rest/auth/send_reset_password"
	"src/presentation/http/router"
)

var Group = router.Group("/auth", func(g *router.GroupRouter) {
	g.Patch("/activate", activate_account.Route)
	g.Post("/activate", resend_activation_code.Route)
	g.Get("/check/email/:email", check_email_available.Route)
	g.Post("/login/credential", login_using_credential.Route)
	g.Patch("/recover", reset_password.Route)
	g.Post("/recover", send_reset_password.Route)
	g.Post("/refresh", refresh_authorization_token.Route)
	g.Post("/register", register_account.Route)
})
