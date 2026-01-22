// domain/entity/subscription.go
package entity

import (
	"encoding/json"
	"src/domain/enum"
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID                    uuid.UUID               `json:"id"`
	TS                    time.Time               `json:"ts"`
	CreatedAt             time.Time               `json:"created_at"`
	DeletedAt             *time.Time              `json:"deleted_at"`
	Status                enum.SubscriptionStatus `json:"status"`
	TrialEndsAt           *time.Time              `json:"trial_ends_at"`
	CurrentPeriodStartAt  *time.Time              `json:"current_period_start_at"`
	CurrentPeriodEndAt    *time.Time              `json:"current_period_end_at"`
	CanceledAt            *time.Time              `json:"canceled_at"`
	Currency              string                  `json:"currency"`
	PriceCents            int                     `json:"price_cents"`
	StripeCustomerKey     string                  `json:"stripe_customer_key"`
	StripeSubscriptionKey string                  `json:"stripe_subscription_key"`
	OrganizationID        uuid.UUID               `json:"organization_id"`
	PlanID                uuid.UUID               `json:"plan_id"`
}

var _ json.Marshaler = (*Subscription)(nil)
var _ json.Unmarshaler = (*Subscription)(nil)

func (e *Subscription) MarshalJSON() ([]byte, error) {
	type Alias Subscription
	return json.Marshal((*Alias)(e))
}

func (e *Subscription) UnmarshalJSON(data []byte) error {
	type Alias Subscription
	return json.Unmarshal(data, (*Alias)(e))
}

