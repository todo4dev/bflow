// domain/tenant/enum/membership.go
package enum

type MembershipRole string

const (
	MembershipRole_ADMIN   MembershipRole = "ADMIN"
	MembershipRole_MANAGER MembershipRole = "MANAGER"
	MembershipRole_VIEWER  MembershipRole = "VIEWER"
	MembershipRole_MEMBER  MembershipRole = "MEMBER"
)
