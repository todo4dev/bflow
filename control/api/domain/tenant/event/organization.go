// domain/tenant/event/organization.go
package event

import (
	"src/domain"

	"github.com/google/uuid"
)

type Organization string

const (
	Organization_CREATED                Organization = "organization.created"
	Organization_UPDATED                Organization = "organization.updated"
	Organization_DELETED                Organization = "organization.deleted"
	OrganizationMembership_ADDED        Organization = "organization.membership_added"
	OrganizationMembership_ROLE_UPDATED Organization = "organization.membership_role_updated"
	OrganizationMembership_REMOVED      Organization = "organization.membership_removed"
	OrganizationInvite_CREATED          Organization = "organization.invite_created"
	OrganizationInvite_RESENT           Organization = "organization.invite_resent"
	OrganizationInvite_CANCELED         Organization = "organization.invite_canceled"
	OrganizationInvite_ACCEPTED         Organization = "organization.invite_accepted"
)

type OrganizationCreatedPayload struct {
	OrganizationID   uuid.UUID `json:"organization_id"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	CreatorAccountID uuid.UUID `json:"creator_account_id"`
}

func OrganizationCreated(
	organizationID uuid.UUID,
	name string,
	slug string,
	creatorAccountID uuid.UUID,
) domain.Event[Organization] {
	return domain.NewEvent(
		Organization_CREATED,
		OrganizationCreatedPayload{
			OrganizationID:   organizationID,
			Name:             name,
			Slug:             slug,
			CreatorAccountID: creatorAccountID,
		},
	)
}

type OrganizationUpdatedPayload struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	Name           string    `json:"name"`
}

func OrganizationUpdated(
	organizationID uuid.UUID,
	name string,
) domain.Event[Organization] {
	return domain.NewEvent(
		Organization_UPDATED,
		OrganizationUpdatedPayload{
			OrganizationID: organizationID,
			Name:           name,
		},
	)
}

type OrganizationDeletedPayload struct {
	OrganizationID uuid.UUID `json:"organization_id"`
}

func OrganizationDeleted(
	organizationID uuid.UUID,
) domain.Event[Organization] {
	return domain.NewEvent(
		Organization_DELETED,
		OrganizationDeletedPayload{
			OrganizationID: organizationID,
		},
	)
}

// Nested entity: Membership
type OrganizationMembershipAddedPayload struct {
	MembershipID   uuid.UUID `json:"membership_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	AccountID      uuid.UUID `json:"account_id"`
	Role           string    `json:"role"`
}

func OrganizationMembershipAdded(
	membershipID uuid.UUID,
	organizationID uuid.UUID,
	accountID uuid.UUID,
	role string,
) domain.Event[Organization] {
	return domain.NewEvent(
		OrganizationMembership_ADDED,
		OrganizationMembershipAddedPayload{
			MembershipID:   membershipID,
			OrganizationID: organizationID,
			AccountID:      accountID,
			Role:           role,
		},
	)
}

type OrganizationMembershipRoleUpdatedPayload struct {
	MembershipID   uuid.UUID `json:"membership_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	AccountID      uuid.UUID `json:"account_id"`
	OldRole        string    `json:"old_role"`
	NewRole        string    `json:"new_role"`
}

func OrganizationMembershipRoleUpdated(
	membershipID uuid.UUID,
	organizationID uuid.UUID,
	accountID uuid.UUID,
	oldRole string,
	newRole string,
) domain.Event[Organization] {
	return domain.NewEvent(
		OrganizationMembership_ROLE_UPDATED,
		OrganizationMembershipRoleUpdatedPayload{
			MembershipID:   membershipID,
			OrganizationID: organizationID,
			AccountID:      accountID,
			OldRole:        oldRole,
			NewRole:        newRole,
		},
	)
}

type OrganizationMembershipRemovedPayload struct {
	MembershipID   uuid.UUID `json:"membership_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	AccountID      uuid.UUID `json:"account_id"`
}

func OrganizationMembershipRemoved(
	membershipID uuid.UUID,
	organizationID uuid.UUID,
	accountID uuid.UUID,
) domain.Event[Organization] {
	return domain.NewEvent(
		OrganizationMembership_REMOVED,
		OrganizationMembershipRemovedPayload{
			MembershipID:   membershipID,
			OrganizationID: organizationID,
			AccountID:      accountID,
		},
	)
}

// Nested entity: Invite
type OrganizationInviteCreatedPayload struct {
	InviteID         uuid.UUID `json:"invite_id"`
	OrganizationID   uuid.UUID `json:"organization_id"`
	Email            string    `json:"email"`
	Role             string    `json:"role"`
	CreatorAccountID uuid.UUID `json:"creator_account_id"`
}

func OrganizationInviteCreated(
	inviteID uuid.UUID,
	organizationID uuid.UUID,
	email string,
	role string,
	creatorAccountID uuid.UUID,
) domain.Event[Organization] {
	return domain.NewEvent(
		OrganizationInvite_CREATED,
		OrganizationInviteCreatedPayload{
			InviteID:         inviteID,
			OrganizationID:   organizationID,
			Email:            email,
			Role:             role,
			CreatorAccountID: creatorAccountID,
		},
	)
}

type OrganizationInviteResentPayload struct {
	InviteID       uuid.UUID `json:"invite_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	Email          string    `json:"email"`
}

func OrganizationInviteResent(
	inviteID uuid.UUID,
	organizationID uuid.UUID,
	email string,
) domain.Event[Organization] {
	return domain.NewEvent(
		OrganizationInvite_RESENT,
		OrganizationInviteResentPayload{
			InviteID:       inviteID,
			OrganizationID: organizationID,
			Email:          email,
		},
	)
}

type OrganizationInviteCanceledPayload struct {
	InviteID       uuid.UUID `json:"invite_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
}

func OrganizationInviteCanceled(
	inviteID uuid.UUID,
	organizationID uuid.UUID,
) domain.Event[Organization] {
	return domain.NewEvent(
		OrganizationInvite_CANCELED,
		OrganizationInviteCanceledPayload{
			InviteID:       inviteID,
			OrganizationID: organizationID,
		},
	)
}

type OrganizationInviteAcceptedPayload struct {
	InviteID       uuid.UUID `json:"invite_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	AccountID      uuid.UUID `json:"account_id"`
}

func OrganizationInviteAccepted(
	inviteID uuid.UUID,
	organizationID uuid.UUID,
	accountID uuid.UUID,
) domain.Event[Organization] {
	return domain.NewEvent(
		OrganizationInvite_ACCEPTED,
		OrganizationInviteAcceptedPayload{
			InviteID:       inviteID,
			OrganizationID: organizationID,
			AccountID:      accountID,
		},
	)
}

var OrganizationMapper = domain.NewEventMapper[Organization]().
	Decoder(Organization_CREATED, domain.JSONDecoder[OrganizationCreatedPayload]()).
	Decoder(Organization_UPDATED, domain.JSONDecoder[OrganizationUpdatedPayload]()).
	Decoder(Organization_DELETED, domain.JSONDecoder[OrganizationDeletedPayload]()).
	Decoder(OrganizationMembership_ADDED, domain.JSONDecoder[OrganizationMembershipAddedPayload]()).
	Decoder(OrganizationMembership_ROLE_UPDATED, domain.JSONDecoder[OrganizationMembershipRoleUpdatedPayload]()).
	Decoder(OrganizationMembership_REMOVED, domain.JSONDecoder[OrganizationMembershipRemovedPayload]()).
	Decoder(OrganizationInvite_CREATED, domain.JSONDecoder[OrganizationInviteCreatedPayload]()).
	Decoder(OrganizationInvite_RESENT, domain.JSONDecoder[OrganizationInviteResentPayload]()).
	Decoder(OrganizationInvite_CANCELED, domain.JSONDecoder[OrganizationInviteCanceledPayload]()).
	Decoder(OrganizationInvite_ACCEPTED, domain.JSONDecoder[OrganizationInviteAcceptedPayload]())
