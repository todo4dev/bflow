// domain/enum/subscription_payment.go
package enum

type SubscriptionPaymentStatus string

const (
	SubscriptionPaymentStatus_PENDING   SubscriptionPaymentStatus = "PENDING"
	SubscriptionPaymentStatus_SUCCEEDED SubscriptionPaymentStatus = "SUCCEEDED"
	SubscriptionPaymentStatus_FAILED    SubscriptionPaymentStatus = "FAILED"
	SubscriptionPaymentStatus_CANCELED  SubscriptionPaymentStatus = "CANCELED"
)

