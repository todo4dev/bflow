// domain/billing/enum/plan.go
package enum

type PlanInterval string

const (
	PlanInterval_HOURLY  PlanInterval = "HOURLY"
	PlanInterval_DAILY   PlanInterval = "DAILY"
	PlanInterval_WEEKLY  PlanInterval = "WEEKLY"
	PlanInterval_MONTHLY PlanInterval = "MONTHLY"
	PlanInterval_YEARLY  PlanInterval = "YEARLY"
)
