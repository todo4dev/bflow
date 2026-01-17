// domain/identity/service/otp.go
package service

import (
	"time"

	"github.com/google/uuid"
)

type Otp interface {
	Generate(accountID uuid.UUID) (code string, expiresAt time.Time, err error)
	Validate(accountID uuid.UUID, code string) (bool, error)
}
