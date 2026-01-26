// infrastructure/database/pgx/repository/subscription_payment.go
package repository

import (
	"context"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type SubscriptionPayment struct {
	client database.Client
}

func NewSubscriptionPayment(client database.Client) *SubscriptionPayment {
	return &SubscriptionPayment{client: client}
}

var _ repository.SubscriptionPayment = (*SubscriptionPayment)(nil)

func (r *SubscriptionPayment) Create(t *entity.SubscriptionPayment) error {
	query := `
		INSERT INTO "subscription_payment" (
			"id", "ts", "created_at", "deleted_at", "status", "currency", "amount_cents",
			"stripe_payment_intent_key", "failure_code", "failure_message", "processed_at",
			"subscription_invoice_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Status, t.Currency, t.AmountCents,
		t.StripePaymentIntentKey, t.FailureCode, t.FailureMessage, t.ProcessedAt,
		t.SubscriptionInvoiceID,
	)
	return err
}

func (r *SubscriptionPayment) Save(t *entity.SubscriptionPayment) error {
	query := `
		UPDATE "subscription_payment" SET
			"ts" = $2, "deleted_at" = $3, "status" = $4, "currency" = $5, "amount_cents" = $6,
			"stripe_payment_intent_key" = $7, "failure_code" = $8,
			"failure_message" = $9, "processed_at" = $10
		WHERE "id" = $1
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.DeletedAt, t.Status, t.Currency, t.AmountCents,
		t.StripePaymentIntentKey, t.FailureCode, t.FailureMessage, t.ProcessedAt,
	)
	return err
}

func (r *SubscriptionPayment) Delete(id uuid.UUID) error {
	query := `DELETE FROM "subscription_payment" WHERE "id" = $1`
	_, err := r.client.Exec(context.Background(), query, id)
	return err
}

func (r *SubscriptionPayment) FindById(id uuid.UUID) (*entity.SubscriptionPayment, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "status", "currency", "amount_cents",
			"stripe_payment_intent_key", "failure_code", "failure_message", "processed_at",
			"subscription_invoice_id"
		FROM "subscription_payment"
		WHERE "id" = $1
	`
	var t entity.SubscriptionPayment
	if err := r.client.QueryRow(context.Background(), query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Status, &t.Currency, &t.AmountCents,
		&t.StripePaymentIntentKey, &t.FailureCode, &t.FailureMessage, &t.ProcessedAt,
		&t.SubscriptionInvoiceID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *SubscriptionPayment) FindByInvoiceId(invoiceId uuid.UUID) ([]*entity.SubscriptionPayment, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "status", "currency", "amount_cents",
			"stripe_payment_intent_key", "failure_code", "failure_message", "processed_at",
			"subscription_invoice_id"
		FROM "subscription_payment"
		WHERE "subscription_invoice_id" = $1
	`
	rows, err := r.client.Query(context.Background(), query, invoiceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.SubscriptionPayment
	for rows.Next() {
		var t entity.SubscriptionPayment
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Status, &t.Currency, &t.AmountCents,
			&t.StripePaymentIntentKey, &t.FailureCode, &t.FailureMessage, &t.ProcessedAt,
			&t.SubscriptionInvoiceID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
