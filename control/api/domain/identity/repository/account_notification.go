// domain/identity/repository/account_notification.go
package repository

import (
	"src/domain/identity/entity"

	"github.com/google/uuid"
)

type AccountNotification interface {
	Create(notification *entity.AccountNotification) error
	Save(notification *entity.AccountNotification) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.AccountNotification, error)
	FindByAccountId(accountId uuid.UUID) ([]*entity.AccountNotification, error)
	FindByClusterId(clusterId uuid.UUID) ([]*entity.AccountNotification, error)
	CountUnreadByAccountId(accountId uuid.UUID) (int64, error)
	MarkAsReadByAccountId(accountId uuid.UUID) error
}
