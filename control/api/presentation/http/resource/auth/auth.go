// presentation/http/resource/auth/auth.go
package auth

import (
	"src/presentation/http/resource/auth/activate_account"
	"src/presentation/http/resource/auth/check_email_available"
	"src/presentation/http/resource/auth/register_account"
	"src/presentation/http/resource/auth/resend_activation_code"
	"src/presentation/http/router"
)

var Group = router.Group("/auth", func(g *router.GroupRouter) {
	g.Get("/check/email/:email", check_email_available.Route)
	g.Post("/register", register_account.Route)
	g.Patch("/activate", activate_account.Route)
	g.Post("/activate", resend_activation_code.Route)
})
