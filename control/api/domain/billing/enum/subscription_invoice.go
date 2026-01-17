// domain/billing/enum/subscription_invoice.go
package enum

type SubscriptionInvoiceStatus string

const (
	SubscriptionInvoiceStatus_OPEN          SubscriptionInvoiceStatus = "OPEN"
	SubscriptionInvoiceStatus_PAID          SubscriptionInvoiceStatus = "PAID"
	SubscriptionInvoiceStatus_VOID          SubscriptionInvoiceStatus = "VOID"
	SubscriptionInvoiceStatus_UNCOLLECTIBLE SubscriptionInvoiceStatus = "UNCOLLECTIBLE"
	SubscriptionInvoiceStatus_DRAFT         SubscriptionInvoiceStatus = "DRAFT"
)
