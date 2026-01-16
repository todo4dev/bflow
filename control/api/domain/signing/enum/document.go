// domain/signing/enum/document.go
package enum

type DocumentKind string

const (
	DocumentKind_AGREEMENT DocumentKind = "AGREEMENT"
	DocumentKind_INVOICE   DocumentKind = "INVOICE"
	DocumentKind_RECEIPT   DocumentKind = "RECEIPT"
	DocumentKind_REPORT    DocumentKind = "REPORT"
	DocumentKind_OTHER     DocumentKind = "OTHER"
)

type DocumentStatus string

const (
	DocumentStatus_ACTIVE   DocumentStatus = "ACTIVE"
	DocumentStatus_REPLACED DocumentStatus = "REPLACED"
	DocumentStatus_REVOKED  DocumentStatus = "REVOKED"
)
