// domain/event/plan.go
package event

import (
	"src/domain"

	"github.com/google/uuid"
)

const (
	Plan_CREATED = "plan.created"
	Plan_UPDATED = "plan.updated"
	Plan_DELETED = "plan.deleted"
)

type PlanCreatedPayload struct {
	PlanID     uuid.UUID `json:"plan_id"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Interval   string    `json:"interval"`
	PriceCents int       `json:"price_cents"`
	Currency   string    `json:"currency"`
}

func PlanCreated(planID uuid.UUID, code, name, interval string, priceCents int, currency string) domain.Event {
	return domain.NewEvent(Plan_CREATED, PlanCreatedPayload{
		PlanID:     planID,
		Code:       code,
		Name:       name,
		Interval:   interval,
		PriceCents: priceCents,
		Currency:   currency,
	})
}

type PlanUpdatedPayload struct {
	PlanID     uuid.UUID `json:"plan_id"`
	Name       string    `json:"name"`
	PriceCents int       `json:"price_cents"`
}

func PlanUpdated(planID uuid.UUID, name string, priceCents int) domain.Event {
	return domain.NewEvent(Plan_UPDATED, PlanUpdatedPayload{
		PlanID:     planID,
		Name:       name,
		PriceCents: priceCents,
	})
}

type PlanDeletedPayload struct {
	PlanID uuid.UUID `json:"plan_id"`
}

func PlanDeleted(planID uuid.UUID) domain.Event {
	return domain.NewEvent(Plan_DELETED, PlanDeletedPayload{
		PlanID: planID,
	})
}

var PlanMapper = domain.NewEventMapper().
	Decoder(Plan_CREATED, domain.JSONDecoder[PlanCreatedPayload]()).
	Decoder(Plan_UPDATED, domain.JSONDecoder[PlanUpdatedPayload]()).
	Decoder(Plan_DELETED, domain.JSONDecoder[PlanDeletedPayload]())
