package check_email_available

import (
	"context"

	"src/domain/identity/issue"
	"src/domain/identity/repository"
)

type Handler struct {
	accountRepo repository.Account
}

func New(accountRepo repository.Account) *Handler {
	return &Handler{accountRepo: accountRepo}
}

func (h *Handler) Handle(ctx context.Context, data *Data) (any, error) {
	exists, err := h.accountRepo.ExistsByEmail(data.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, issue.EmailInUse{Email: data.Email}
	}
	return nil, nil
}
