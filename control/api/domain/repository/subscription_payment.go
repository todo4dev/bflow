// domain/repository/subscription_payment.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type SubscriptionPayment interface {
	Create(payment *entity.SubscriptionPayment) error
	Save(payment *entity.SubscriptionPayment) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.SubscriptionPayment, error)
	FindByInvoiceId(invoiceId uuid.UUID) ([]*entity.SubscriptionPayment, error)
}
