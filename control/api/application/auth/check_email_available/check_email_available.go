// application/auth/check_email_available/check_email_available.go
package check_email_available

import (
	"context"

	"src/domain/issue"
	"src/domain/repository"

	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/validate"
)

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.Email).Text().Email().Required()
})

type Data struct {
	Email string `json:"email"`
}

type Handler struct {
	repositoryAccount repository.Account
}

func New(
	repositoryAccount repository.Account,
) *Handler {
	return &Handler{
		repositoryAccount: repositoryAccount,
	}
}

func (h *Handler) checkEmailInUse(email string) error {
	exists, err := h.repositoryAccount.ExistsByEmail(email)
	if err != nil {
		return err
	}
	if exists {
		return &issue.AccountEmailInUse{}
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, data *Data) error {
	if _, err := dataSchema.Validate(data); err != nil {
		return err
	}
	return h.checkEmailInUse(data.Email)
}

func init() {
	data := Data{
		Email: "john.doe@email.com"}
	meta.Describe(&data, meta.Description("Data for checking email availability"),
		meta.Field(&data.Email, meta.Description("Email to check")),
		meta.Example(data))
	meta.Describe(&Handler{}, meta.Description("Handler for checking email availability"),
		meta.Throws[*issue.AccountEmailInUse]())
}
