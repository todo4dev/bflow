// infrastructure/database/pgx/repository/subscription_invoice.go
package repository

import (
	"context"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type SubscriptionInvoice struct {
	client database.Client
}

func NewSubscriptionInvoice(client database.Client) *SubscriptionInvoice {
	return &SubscriptionInvoice{client: client}
}

var _ repository.SubscriptionInvoice = (*SubscriptionInvoice)(nil)

func (r *SubscriptionInvoice) Create(t *entity.SubscriptionInvoice) error {
	query := `
		INSERT INTO "subscription_invoice" (
			"id", "ts", "created_at", "deleted_at", "status", "currency",
			"total_cents", "tax_cents", "discount_cents", "due_at", "paid_at",
			"stripe_invoice_key", "subscription_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Status, t.Currency,
		t.TotalCents, t.TaxCents, t.DiscountCents, t.DueAt, t.PaidAt,
		t.StripeInvoiceKey, t.SubscriptionID,
	)
	return err
}

func (r *SubscriptionInvoice) Save(t *entity.SubscriptionInvoice) error {
	query := `
		UPDATE "subscription_invoice" SET
			"ts" = $2, "deleted_at" = $3, "status" = $4, "currency" = $5,
			"total_cents" = $6, "tax_cents" = $7, "discount_cents" = $8,
			"due_at" = $9, "paid_at" = $10, "stripe_invoice_key" = $11
		WHERE "id" = $1
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.DeletedAt, t.Status, t.Currency,
		t.TotalCents, t.TaxCents, t.DiscountCents, t.DueAt, t.PaidAt,
		t.StripeInvoiceKey,
	)
	return err
}

func (r *SubscriptionInvoice) Delete(id uuid.UUID) error {
	query := `DELETE FROM "subscription_invoice" WHERE "id" = $1`
	_, err := r.client.Exec(context.Background(), query, id)
	return err
}

func (r *SubscriptionInvoice) FindById(id uuid.UUID) (*entity.SubscriptionInvoice, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "status", "currency",
			"total_cents", "tax_cents", "discount_cents", "due_at", "paid_at",
			"stripe_invoice_key", "subscription_id"
		FROM "subscription_invoice"
		WHERE "id" = $1
	`
	var t entity.SubscriptionInvoice
	if err := r.client.QueryRow(context.Background(), query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Status, &t.Currency,
		&t.TotalCents, &t.TaxCents, &t.DiscountCents, &t.DueAt, &t.PaidAt,
		&t.StripeInvoiceKey, &t.SubscriptionID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *SubscriptionInvoice) FindBySubscriptionId(subscriptionId uuid.UUID) ([]*entity.SubscriptionInvoice, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "status", "currency",
			"total_cents", "tax_cents", "discount_cents", "due_at", "paid_at",
			"stripe_invoice_key", "subscription_id"
		FROM "subscription_invoice"
		WHERE "subscription_id" = $1
	`
	rows, err := r.client.Query(context.Background(), query, subscriptionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.SubscriptionInvoice
	for rows.Next() {
		var t entity.SubscriptionInvoice
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Status, &t.Currency,
			&t.TotalCents, &t.TaxCents, &t.DiscountCents, &t.DueAt, &t.PaidAt,
			&t.StripeInvoiceKey, &t.SubscriptionID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
