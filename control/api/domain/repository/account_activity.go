// domain/repository/account_activity.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type AccountActivity interface {
	Create(activity *entity.AccountActivity) error
	FindById(id uuid.UUID) (*entity.AccountActivity, error)
	FindByAccountId(accountId uuid.UUID) ([]*entity.AccountActivity, error)
	FindByClusterId(clusterId uuid.UUID) ([]*entity.AccountActivity, error)
}

