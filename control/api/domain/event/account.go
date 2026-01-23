// domain/event/account.go
package event

import (
	"src/domain"

	"github.com/google/uuid"
)

const (
	Account_REGISTERED         = "account.registered"
	Account_ROLE_SET           = "account.role_set"
	Account_ACTIVATED          = "account.activated"
	Account_CREDENTIAL_CHANGED = "account.credential_changed"
)

type AccountRegisteredPayload struct {
	Email string `json:"email"`
}

func AccountRegistered(email string) domain.Event {
	return domain.NewEvent(Account_REGISTERED, AccountRegisteredPayload{
		Email: email,
	})
}

type AccountRoleSetPayload struct {
	AccountID uuid.UUID `json:"account_id"`
	Role      string    `json:"role"`
}

func AccountRoleSet(accountID uuid.UUID, role string) domain.Event {
	return domain.NewEvent(Account_ROLE_SET, AccountRoleSetPayload{
		AccountID: accountID,
		Role:      role,
	})
}

type AccountActivatedPayload struct {
	AccountID uuid.UUID `json:"account_id"`
}

func AccountActivated(accountID uuid.UUID) domain.Event {
	return domain.NewEvent(Account_ACTIVATED, AccountActivatedPayload{
		AccountID: accountID,
	})
}

type AccountCredentialChangedPayload struct {
	AccountID uuid.UUID `json:"account_id"`
}

func AccountCredentialChanged(accountID uuid.UUID) domain.Event {
	return domain.NewEvent(Account_CREDENTIAL_CHANGED, AccountCredentialChangedPayload{
		AccountID: accountID,
	})
}

var AccountMapper = domain.NewEventMapper().
	Decoder(Account_REGISTERED, domain.JSONDecoder[AccountRegisteredPayload]()).
	Decoder(Account_ROLE_SET, domain.JSONDecoder[AccountRoleSetPayload]()).
	Decoder(Account_ACTIVATED, domain.JSONDecoder[AccountActivatedPayload]()).
	Decoder(Account_CREDENTIAL_CHANGED, domain.JSONDecoder[AccountCredentialChangedPayload]())
