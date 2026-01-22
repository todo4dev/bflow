// domain/repository/subscription_invoice.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type SubscriptionInvoice interface {
	Create(invoice *entity.SubscriptionInvoice) error
	Save(invoice *entity.SubscriptionInvoice) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.SubscriptionInvoice, error)
	FindBySubscriptionId(subscriptionId uuid.UUID) ([]*entity.SubscriptionInvoice, error)
}
