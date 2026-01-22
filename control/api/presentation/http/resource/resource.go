// presentation/http/resource/resource.go
package resource

import (
	"src/presentation/http/resource/auth"
	"src/presentation/http/resource/system"
	"src/presentation/http/router"
)

var Routes = []router.GroupDefinition{
	auth.Group,
	system.Group,
}
