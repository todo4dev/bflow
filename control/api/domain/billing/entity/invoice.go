// domain/billing/entity/invoice.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// invoice

type Invoice_StatusEnum string

const (
	Invoice_StatusOpen          Invoice_StatusEnum = "open"
	Invoice_StatusPaid          Invoice_StatusEnum = "paid"
	Invoice_StatusVoid          Invoice_StatusEnum = "void"
	Invoice_StatusUncollectible Invoice_StatusEnum = "uncollectible"
	Invoice_StatusDraft         Invoice_StatusEnum = "draft"
)

type InvoiceEntity struct {
	ID               uuid.UUID          `json:"id"`
	TS               time.Time          `json:"ts"`
	CreatedAt        time.Time          `json:"created_at"`
	DeletedAt        *time.Time         `json:"deleted_at"`
	Status           Invoice_StatusEnum `json:"status"`
	Currency         string             `json:"currency"`
	TotalCents       int                `json:"total_cents"`
	TaxCents         int                `json:"tax_cents"`
	DiscountCents    int                `json:"discount_cents"`
	DueAt            *time.Time         `json:"due_at"`
	PaidAt           *time.Time         `json:"paid_at"`
	StripeInvoiceKey string             `json:"stripe_invoice_key"`
	SubscriptionID   uuid.UUID          `json:"subscription_id"`
}

var _ json.Marshaler = (*InvoiceEntity)(nil)
var _ json.Unmarshaler = (*InvoiceEntity)(nil)

func (e *InvoiceEntity) MarshalJSON() ([]byte, error) {
	type Alias InvoiceEntity
	return json.Marshal((*Alias)(e))
}

func (e *InvoiceEntity) UnmarshalJSON(data []byte) error {
	type Alias InvoiceEntity
	return json.Unmarshal(data, (*Alias)(e))
}
