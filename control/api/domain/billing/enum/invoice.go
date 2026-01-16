// domain/billing/enum/invoice.go
package enum

type InvoiceStatus string

const (
	InvoiceStatus_OPEN          InvoiceStatus = "OPEN"
	InvoiceStatus_PAID          InvoiceStatus = "PAID"
	InvoiceStatus_VOID          InvoiceStatus = "VOID"
	InvoiceStatus_UNCOLLECTIBLE InvoiceStatus = "UNCOLLECTIBLE"
	InvoiceStatus_DRAFT         InvoiceStatus = "DRAFT"
)
