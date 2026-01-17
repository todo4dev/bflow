// domain/tenant/enum/organization_membership.go
package enum

type OrganizationMembershipRole string

const (
	OrganizationMembershipRole_ADMIN   OrganizationMembershipRole = "ADMIN"
	OrganizationMembershipRole_MANAGER OrganizationMembershipRole = "MANAGER"
	OrganizationMembershipRole_VIEWER  OrganizationMembershipRole = "VIEWER"
	OrganizationMembershipRole_MEMBER  OrganizationMembershipRole = "MEMBER"
)
