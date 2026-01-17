// domain/tenant/enum/organization_invite.go
package enum

type OrganizationInviteStatus string

const (
	OrganizationInviteStatus_PENDING  OrganizationInviteStatus = "PENDING"
	OrganizationInviteStatus_ACCEPTED OrganizationInviteStatus = "ACCEPTED"
	OrganizationInviteStatus_EXPIRED  OrganizationInviteStatus = "EXPIRED"
	OrganizationInviteStatus_REVOKED  OrganizationInviteStatus = "REVOKED"
)
