// domain/identity/event/account.go
package event

import (
	"src/domain"

	"github.com/google/uuid"
)

type Account string

const (
	Account_REGISTERED         Account = "account.registered"
	Account_ROLE_SET           Account = "account.role_set"
	Account_ACTIVATED          Account = "account.activated"
	Account_CREDENTIAL_CHANGED Account = "account.credential_changed"
)

type AccountRegisteredPayload struct {
	Email string `json:"email"`
}

func AccountRegistered(
	email string,
) domain.Event[Account] {
	return domain.NewEvent(
		Account_REGISTERED,
		AccountRegisteredPayload{
			Email: email,
		},
	)
}

type AccountRoleSetPayload struct {
	AccountID uuid.UUID `json:"account_id"`
	Role      string    `json:"role"`
}

func AccountRoleSet(
	accountID uuid.UUID,
	role string,
) domain.Event[Account] {
	return domain.NewEvent(
		Account_ROLE_SET,
		AccountRoleSetPayload{
			AccountID: accountID,
			Role:      role,
		},
	)
}

type AccountActivatedPayload struct {
	AccountID uuid.UUID `json:"account_id"`
}

func AccountActivated(
	accountID uuid.UUID,
) domain.Event[Account] {
	return domain.NewEvent(
		Account_ACTIVATED,
		AccountActivatedPayload{
			AccountID: accountID,
		},
	)
}

type AccountCredentialChangedPayload struct {
	AccountID uuid.UUID `json:"account_id"`
}

func AccountCredentialChanged(
	accountID uuid.UUID,
) domain.Event[Account] {
	return domain.NewEvent(
		Account_CREDENTIAL_CHANGED,
		AccountCredentialChangedPayload{
			AccountID: accountID,
		},
	)
}

var AccountMapper = domain.NewEventMapper[Account]().
	Decoder(Account_REGISTERED, domain.JSONDecoder[AccountRegisteredPayload]()).
	Decoder(Account_ROLE_SET, domain.JSONDecoder[AccountRoleSetPayload]()).
	Decoder(Account_ACTIVATED, domain.JSONDecoder[AccountActivatedPayload]()).
	Decoder(Account_CREDENTIAL_CHANGED, domain.JSONDecoder[AccountCredentialChangedPayload]())
