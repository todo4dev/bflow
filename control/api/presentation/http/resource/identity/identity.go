package identity

import (
	"src/presentation/http/resource/identity/check_email_available"
	"src/presentation/http/router"
)

var Group = router.Group("/identity", func(g *router.GroupRouter) {
	g.Get("/check/email/:email", check_email_available.Route)
})
