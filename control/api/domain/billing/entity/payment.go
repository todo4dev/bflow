// domain/billing/entity/payment.go
// domain/billing/entity/invoice.go
package entity

import (
	"encoding/json"
	"src/domain/billing/enum"
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID                     uuid.UUID          `json:"id"`
	TS                     time.Time          `json:"ts"`
	CreatedAt              time.Time          `json:"created_at"`
	DeletedAt              *time.Time         `json:"deleted_at"`
	Status                 enum.PaymentStatus `json:"status"`
	Currency               string             `json:"currency"`
	AmountCents            int                `json:"amount_cents"`
	StripePaymentIntentKey string             `json:"stripe_payment_intent_key"`
	FailureCode            *string            `json:"failure_code"`
	FailureMessage         *string            `json:"failure_message"`
	ProcessedAt            *time.Time         `json:"processed_at"`
	InvoiceID              uuid.UUID          `json:"invoice_id"`
}

var _ json.Marshaler = (*Payment)(nil)
var _ json.Unmarshaler = (*Payment)(nil)

func (e *Payment) MarshalJSON() ([]byte, error) {
	type Alias Payment
	return json.Marshal((*Alias)(e))
}

func (e *Payment) UnmarshalJSON(data []byte) error {
	type Alias Payment
	return json.Unmarshal(data, (*Alias)(e))
}
