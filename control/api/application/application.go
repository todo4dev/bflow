// application/application.go
package application

import (
	"src/application/auth"
	"src/application/health"
)

func Provide() {
	auth.Provide()
	health.Provide()
}
