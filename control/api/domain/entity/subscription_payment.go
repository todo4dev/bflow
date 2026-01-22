// domain/entity/subscription_payment.go
package entity

import (
	"encoding/json"
	"src/domain/enum"
	"time"

	"github.com/google/uuid"
)

type SubscriptionPayment struct {
	ID                     uuid.UUID                      `json:"id"`
	TS                     time.Time                      `json:"ts"`
	CreatedAt              time.Time                      `json:"created_at"`
	DeletedAt              *time.Time                     `json:"deleted_at"`
	Status                 enum.SubscriptionPaymentStatus `json:"status"`
	Currency               string                         `json:"currency"`
	AmountCents            int                            `json:"amount_cents"`
	StripePaymentIntentKey string                         `json:"stripe_payment_intent_key"`
	FailureCode            *string                        `json:"failure_code"`
	FailureMessage         *string                        `json:"failure_message"`
	ProcessedAt            *time.Time                     `json:"processed_at"`
	SubscriptionInvoiceID  uuid.UUID                      `json:"subscription_invoice_id"`
}

var _ json.Marshaler = (*SubscriptionPayment)(nil)
var _ json.Unmarshaler = (*SubscriptionPayment)(nil)

func (e *SubscriptionPayment) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionPayment
	return json.Marshal((*Alias)(e))
}

func (e *SubscriptionPayment) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionPayment
	return json.Unmarshal(data, (*Alias)(e))
}

