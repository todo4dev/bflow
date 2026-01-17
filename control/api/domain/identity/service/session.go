// domain/identity/service/session.go
package service

import (
	"time"

	"github.com/google/uuid"
)

type Session interface {
	Create(accountID uuid.UUID, maxTTL time.Duration) (sessionID string, err error)
	Get(sessionID string) (maxTTL time.Duration, err error)
	Revoke(sessionID string) error
	RevokeAll(accountID uuid.UUID) error
}
