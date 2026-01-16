// domain/tenant/enum/invite.go
package enum

type InviteStatus string

const (
	InviteStatus_PENDING  InviteStatus = "PENDING"
	InviteStatus_ACCEPTED InviteStatus = "ACCEPTED"
	InviteStatus_EXPIRED  InviteStatus = "EXPIRED"
	InviteStatus_REVOKED  InviteStatus = "REVOKED"
)
