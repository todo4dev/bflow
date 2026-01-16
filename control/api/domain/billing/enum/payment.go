// domain/billing/enum/payment.go
package enum

type PaymentStatus string

const (
	PaymentStatus_PENDING   PaymentStatus = "PENDING"
	PaymentStatus_SUCCEEDED PaymentStatus = "SUCCEEDED"
	PaymentStatus_FAILED    PaymentStatus = "FAILED"
	PaymentStatus_CANCELED  PaymentStatus = "CANCELED"
)
