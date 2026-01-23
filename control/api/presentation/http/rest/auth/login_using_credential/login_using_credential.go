// presentation/http/rest/auth/login_using_credential/login_using_credential.go
package login_using_credential

import (
	usecase "src/application/auth/login_using_credential"
	"src/domain/issue"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
	"github.com/leandroluk/gox/validate"
)

func handler(c *fiber.Ctx) error {
	var data usecase.Data
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	result, err := di.Resolve[*usecase.Handler]().Handle(c.Context(), &data)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func operation(o *oas.Operation) {
	o.Tags("Auth").Summary("Login using credential").
		Description("Authenticates a user by verifying their email and password.")
	router.BodyJson(o, router.SchemaAs[usecase.Data]())
	router.ResponseStatus(o, fiber.StatusOK, "Login successful", router.SchemaAs[usecase.Result]())
	router.ResponseIssueAs[*issue.AccountInvalidCredentials](o, fiber.StatusUnauthorized)
	router.ResponseIssueAs[*issue.AccountNotVerified](o, fiber.StatusForbidden)
	router.ResponseIssueAs[*issue.AccountDeactivated](o, fiber.StatusNotAcceptable)
	router.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusUnprocessableEntity)
}

var Route = router.
	Route(handler).
	Operation(operation)
