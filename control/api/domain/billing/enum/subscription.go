// domain/billing/enum/subscription.go

package enum

type SubscriptionStatus string

const (
	SubscriptionStatus_TRIALING SubscriptionStatus = "TRIALING"
	SubscriptionStatus_ACTIVE   SubscriptionStatus = "ACTIVE"
	SubscriptionStatus_PAST_DUE SubscriptionStatus = "PAST_DUE"
	SubscriptionStatus_CANCELED SubscriptionStatus = "CANCELED"
)
