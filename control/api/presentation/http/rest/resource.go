// presentation/http/rest/resource.go
package resource

import (
	"src/presentation/http/rest/auth"
	"src/presentation/http/rest/system"
	"src/presentation/http/router"
)

var Routes = []router.GroupDefinition{
	auth.Group,
	system.Group,
}
