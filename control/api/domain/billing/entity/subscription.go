// subscription

package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Subscription_StatusEnum string

const (
	Subscription_StatusTrialing Subscription_StatusEnum = "trialing"
	Subscription_StatusActive   Subscription_StatusEnum = "active"
	Subscription_StatusPastDue  Subscription_StatusEnum = "past_due"
	Subscription_StatusCanceled Subscription_StatusEnum = "canceled"
)

type SubscriptionEntity struct {
	ID                    uuid.UUID               `json:"id"`
	TS                    time.Time               `json:"ts"`
	CreatedAt             time.Time               `json:"created_at"`
	DeletedAt             *time.Time              `json:"deleted_at"`
	Status                Subscription_StatusEnum `json:"status"`
	Interval              string                  `json:"interval"`
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

var _ json.Marshaler = (*SubscriptionEntity)(nil)
var _ json.Unmarshaler = (*SubscriptionEntity)(nil)

func (e *SubscriptionEntity) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionEntity
	return json.Marshal((*Alias)(e))
}

func (e *SubscriptionEntity) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionEntity
	return json.Unmarshal(data, (*Alias)(e))
}
