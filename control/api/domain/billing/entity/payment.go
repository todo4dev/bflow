// control/api/domain/billing/entity/invoice.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Payment_StatusEnum string

const (
	Payment_StatusPending   Payment_StatusEnum = "pending"
	Payment_StatusSucceeded Payment_StatusEnum = "succeeded"
	Payment_StatusFailed    Payment_StatusEnum = "failed"
	Payment_StatusCanceled  Payment_StatusEnum = "canceled"
)

type PaymentEntity struct {
	ID                     uuid.UUID          `json:"id"`
	TS                     time.Time          `json:"ts"`
	CreatedAt              time.Time          `json:"created_at"`
	DeletedAt              *time.Time         `json:"deleted_at"`
	Status                 Payment_StatusEnum `json:"status"`
	Currency               string             `json:"currency"`
	AmountCents            int                `json:"amount_cents"`
	StripePaymentIntentKey string             `json:"stripe_payment_intent_key"`
	FailureCode            *string            `json:"failure_code"`
	FailureMessage         *string            `json:"failure_message"`
	ProcessedAt            *time.Time         `json:"processed_at"`
	InvoiceID              uuid.UUID          `json:"invoice_id"`
}

var _ json.Marshaler = (*PaymentEntity)(nil)
var _ json.Unmarshaler = (*PaymentEntity)(nil)

func (e *PaymentEntity) MarshalJSON() ([]byte, error) {
	type Alias PaymentEntity
	return json.Marshal((*Alias)(e))
}

func (e *PaymentEntity) UnmarshalJSON(data []byte) error {
	type Alias PaymentEntity
	return json.Unmarshal(data, (*Alias)(e))
}
