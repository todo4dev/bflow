// domain/entity/subscription_invoice.go
package entity

import (
	"encoding/json"
	"src/domain/enum"
	"time"

	"github.com/google/uuid"
)

type SubscriptionInvoice struct {
	ID               uuid.UUID                      `json:"id"`
	TS               time.Time                      `json:"ts"`
	CreatedAt        time.Time                      `json:"created_at"`
	DeletedAt        *time.Time                     `json:"deleted_at"`
	Status           enum.SubscriptionInvoiceStatus `json:"status"`
	Currency         string                         `json:"currency"`
	TotalCents       int                            `json:"total_cents"`
	TaxCents         int                            `json:"tax_cents"`
	DiscountCents    int                            `json:"discount_cents"`
	DueAt            *time.Time                     `json:"due_at"`
	PaidAt           *time.Time                     `json:"paid_at"`
	StripeInvoiceKey string                         `json:"stripe_invoice_key"`
	SubscriptionID   uuid.UUID                      `json:"subscription_id"`
}

var _ json.Marshaler = (*SubscriptionInvoice)(nil)
var _ json.Unmarshaler = (*SubscriptionInvoice)(nil)

func (e *SubscriptionInvoice) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionInvoice
	return json.Marshal((*Alias)(e))
}

func (e *SubscriptionInvoice) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionInvoice
	return json.Unmarshal(data, (*Alias)(e))
}

