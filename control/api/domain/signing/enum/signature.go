// domain/signing/enum/signature.go
package enum

type SignatureState string

const (
	SignatureState_PENDING  SignatureState = "PENDING"
	SignatureState_SIGNED   SignatureState = "SIGNED"
	SignatureState_FAILED   SignatureState = "FAILED"
	SignatureState_CANCELED SignatureState = "CANCELED"
)
