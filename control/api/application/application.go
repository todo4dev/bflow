package application

import (
	"src/application/billing"
	"src/application/deployment"
	"src/application/identity"
	"src/application/signing"
	"src/application/system"
	"src/application/tenant"
)

func Provide() {
	billing.Provide()
	deployment.Provide()
	identity.Provide()
	signing.Provide()
	system.Provide()
	tenant.Provide()
}
