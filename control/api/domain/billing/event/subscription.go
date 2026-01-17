// domain/billing/event/subscription.go
package event

import (
	"src/domain"
	"time"

	"github.com/google/uuid"
)

type Subscription string

const (
	Subscription_CREATED          Subscription = "subscription.created"
	Subscription_PLAN_CHANGED     Subscription = "subscription.plan_changed"
	Subscription_CANCELED         Subscription = "subscription.canceled"
	Subscription_RESUMED          Subscription = "subscription.resumed"
	SubscriptionInvoice_GENERATED Subscription = "subscription.invoice_generated"
	SubscriptionPayment_PROCESSED Subscription = "subscription.payment_processed"
	SubscriptionPayment_FAILED    Subscription = "subscription.payment_failed"
)

type SubscriptionCreatedPayload struct {
	SubscriptionID uuid.UUID  `json:"subscription_id"`
	OrganizationID uuid.UUID  `json:"organization_id"`
	PlanID         uuid.UUID  `json:"plan_id"`
	TrialEndsAt    *time.Time `json:"trial_ends_at"`
}

func SubscriptionCreated(
	subscriptionID uuid.UUID,
	organizationID uuid.UUID,
	planID uuid.UUID,
	trialEndsAt *time.Time,
) domain.Event[Subscription] {
	return domain.NewEvent(
		Subscription_CREATED,
		SubscriptionCreatedPayload{
			SubscriptionID: subscriptionID,
			OrganizationID: organizationID,
			PlanID:         planID,
			TrialEndsAt:    trialEndsAt,
		},
	)
}

type SubscriptionPlanChangedPayload struct {
	SubscriptionID uuid.UUID `json:"subscription_id"`
	OldPlanID      uuid.UUID `json:"old_plan_id"`
	NewPlanID      uuid.UUID `json:"new_plan_id"`
}

func SubscriptionPlanChanged(
	subscriptionID uuid.UUID,
	oldPlanID uuid.UUID,
	newPlanID uuid.UUID,
) domain.Event[Subscription] {
	return domain.NewEvent(
		Subscription_PLAN_CHANGED,
		SubscriptionPlanChangedPayload{
			SubscriptionID: subscriptionID,
			OldPlanID:      oldPlanID,
			NewPlanID:      newPlanID,
		},
	)
}

type SubscriptionCanceledPayload struct {
	SubscriptionID uuid.UUID `json:"subscription_id"`
}

func SubscriptionCanceled(
	subscriptionID uuid.UUID,
) domain.Event[Subscription] {
	return domain.NewEvent(
		Subscription_CANCELED,
		SubscriptionCanceledPayload{
			SubscriptionID: subscriptionID,
		},
	)
}

type SubscriptionResumedPayload struct {
	SubscriptionID uuid.UUID `json:"subscription_id"`
}

func SubscriptionResumed(
	subscriptionID uuid.UUID,
) domain.Event[Subscription] {
	return domain.NewEvent(
		Subscription_RESUMED,
		SubscriptionResumedPayload{
			SubscriptionID: subscriptionID,
		},
	)
}

// Nested entity: Invoice
type SubscriptionInvoiceGeneratedPayload struct {
	InvoiceID      uuid.UUID `json:"invoice_id"`
	SubscriptionID uuid.UUID `json:"subscription_id"`
	TotalCents     int       `json:"total_cents"`
	Currency       string    `json:"currency"`
	DueAt          time.Time `json:"due_at"`
}

func SubscriptionInvoiceGenerated(
	invoiceID uuid.UUID,
	subscriptionID uuid.UUID,
	totalCents int,
	currency string,
	dueAt time.Time,
) domain.Event[Subscription] {
	return domain.NewEvent(
		SubscriptionInvoice_GENERATED,
		SubscriptionInvoiceGeneratedPayload{
			InvoiceID:      invoiceID,
			SubscriptionID: subscriptionID,
			TotalCents:     totalCents,
			Currency:       currency,
			DueAt:          dueAt,
		},
	)
}

// Nested entity: Payment
type SubscriptionPaymentProcessedPayload struct {
	PaymentID   uuid.UUID `json:"payment_id"`
	InvoiceID   uuid.UUID `json:"invoice_id"`
	AmountCents int       `json:"amount_cents"`
	Currency    string    `json:"currency"`
}

func SubscriptionPaymentProcessed(
	paymentID uuid.UUID,
	invoiceID uuid.UUID,
	amountCents int,
	currency string,
) domain.Event[Subscription] {
	return domain.NewEvent(
		SubscriptionPayment_PROCESSED,
		SubscriptionPaymentProcessedPayload{
			PaymentID:   paymentID,
			InvoiceID:   invoiceID,
			AmountCents: amountCents,
			Currency:    currency,
		},
	)
}

type SubscriptionPaymentFailedPayload struct {
	PaymentID      uuid.UUID `json:"payment_id"`
	InvoiceID      uuid.UUID `json:"invoice_id"`
	FailureCode    string    `json:"failure_code"`
	FailureMessage string    `json:"failure_message"`
}

func SubscriptionPaymentFailed(
	paymentID uuid.UUID,
	invoiceID uuid.UUID,
	failureCode string,
	failureMessage string,
) domain.Event[Subscription] {
	return domain.NewEvent(
		SubscriptionPayment_FAILED,
		SubscriptionPaymentFailedPayload{
			PaymentID:      paymentID,
			InvoiceID:      invoiceID,
			FailureCode:    failureCode,
			FailureMessage: failureMessage,
		},
	)
}

var SubscriptionMapper = domain.NewEventMapper[Subscription]().
	Decoder(Subscription_CREATED, domain.JSONDecoder[SubscriptionCreatedPayload]()).
	Decoder(Subscription_PLAN_CHANGED, domain.JSONDecoder[SubscriptionPlanChangedPayload]()).
	Decoder(Subscription_CANCELED, domain.JSONDecoder[SubscriptionCanceledPayload]()).
	Decoder(Subscription_RESUMED, domain.JSONDecoder[SubscriptionResumedPayload]()).
	Decoder(SubscriptionInvoice_GENERATED, domain.JSONDecoder[SubscriptionInvoiceGeneratedPayload]()).
	Decoder(SubscriptionPayment_PROCESSED, domain.JSONDecoder[SubscriptionPaymentProcessedPayload]()).
	Decoder(SubscriptionPayment_FAILED, domain.JSONDecoder[SubscriptionPaymentFailedPayload]())
