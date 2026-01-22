// application/application.go
package application

import (
	"src/application/auth"
	"src/application/system"
)

func Provide() {
	auth.Provide()
	system.Provide()
}
