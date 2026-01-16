// domain/identity/event/account.go
package event

import (
	"src/domain/shared"

	"github.com/google/uuid"
)

type AccountEventEnum string

const (
	AccountEvent_Registered        AccountEventEnum = "AccountEvent_Registered"
	AccountEvent_RoleSet           AccountEventEnum = "AccountEvent_RoleSet"
	AccountEvent_Activated         AccountEventEnum = "AccountEvent_Activated"
	AccountEvent_CredentialChanged AccountEventEnum = "AccountEvent_CredentialChanged"
)

type AccountRegisteredEventPayload struct {
	Email string `json:"email"`
}

func AccountRegisteredEvent(email string) shared.Event[AccountEventEnum] {
	return shared.NewEvent(
		AccountEvent_Registered,
		AccountRegisteredEventPayload{Email: email},
	)
}

type AccountRoleSetEventPayload struct {
	AccountID uuid.UUID `json:"account_id"`
	Role      string    `json:"role"`
}

func AccountRoleSetEvent(accountID uuid.UUID, role string) shared.Event[AccountEventEnum] {
	return shared.NewEvent(
		AccountEvent_RoleSet,
		AccountRoleSetEventPayload{AccountID: accountID, Role: role},
	)
}

type AccountActivatedEventPayload struct {
	AccountID uuid.UUID `json:"account_id"`
}

func AccountActivatedEvent(accountID uuid.UUID) shared.Event[AccountEventEnum] {
	return shared.NewEvent(
		AccountEvent_Activated,
		AccountActivatedEventPayload{AccountID: accountID},
	)
}

type AccountCredentialChangedEventPayload struct {
	AccountID uuid.UUID `json:"account_id"`
}

func AccountCredentialChangedEvent(accountID uuid.UUID) shared.Event[AccountEventEnum] {
	return shared.NewEvent(
		AccountEvent_CredentialChanged,
		AccountCredentialChangedEventPayload{AccountID: accountID},
	)
}

var AccountEventMapper = shared.NewEventMapper[AccountEventEnum]().
	Decoder(AccountEvent_Registered, shared.JSONDecoder[AccountRegisteredEventPayload]()).
	Decoder(AccountEvent_RoleSet, shared.JSONDecoder[AccountRoleSetEventPayload]()).
	Decoder(AccountEvent_Activated, shared.JSONDecoder[AccountActivatedEventPayload]()).
	Decoder(AccountEvent_CredentialChanged, shared.JSONDecoder[AccountCredentialChangedEventPayload]())
