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
	accountRepo repository.Account
}

func New(
	accountRepo repository.Account,
) *Handler {
	return &Handler{
		accountRepo: accountRepo,
	}
}

func (h *Handler) Handle(ctx context.Context, data *Data) (any, error) {
	_, err := dataSchema.Validate(data)
	if err != nil {
		return nil, err
	}
	exists, err := h.accountRepo.ExistsByEmail(data.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &issue.AccountEmailInUse{}
	}
	return nil, nil
}

func init() {
	data := Data{
		Email: "john.doe@email.com"}
	meta.Describe(&data, meta.Description("Data for checking email availability"),
		meta.Field(&data.Email, meta.Description("Email to check")),
		meta.Example(data))
	meta.Describe(&Handler{}, meta.Description("Handler for checking email availability"),
		meta.Throws[issue.AccountEmailInUse](issue.AccountEmailInUse_MESSAGE))
}
