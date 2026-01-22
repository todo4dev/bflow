// domain/enum/document_signature.go
package enum

type DocumentSignatureState string

const (
	DocumentSignatureState_PENDING  DocumentSignatureState = "PENDING"
	DocumentSignatureState_SIGNED   DocumentSignatureState = "SIGNED"
	DocumentSignatureState_FAILED   DocumentSignatureState = "FAILED"
	DocumentSignatureState_CANCELED DocumentSignatureState = "CANCELED"
)

