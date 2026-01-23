// presentation/http/rest/auth/refresh_authorization_token/refresh_authorization_token.go
package refresh_authorization_token

import (
	usecase "src/application/auth/refresh_authorization_token"
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
	o.Tags("Auth").Summary("Refresh Authorization Token").
		Description("Refresh the authorization token extending the session expiration.")
	router.BodyJson(o, router.SchemaAs[usecase.Data]())
	router.ResponseStatus(o, fiber.StatusOK, "Refresh successful", router.SchemaAs[usecase.Result]())
	router.ResponseIssueAs[*issue.AccountInvalidToken](o, fiber.StatusUnauthorized)
	router.ResponseIssueAs[*issue.AccountSessionExpired](o, fiber.StatusUnauthorized)
	router.ResponseIssueAs[*issue.AccountDeactivated](o, fiber.StatusForbidden)
	router.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusUnprocessableEntity)
}

var Route = router.
	Route(handler).
	Operation(operation)
