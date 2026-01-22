// infrastructure/database/pgx/repository/subscription.go
package repository

import (
	"context"
	"time"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Subscription struct {
	client database.Client
}

func NewSubscription(client database.Client) *Subscription {
	return &Subscription{client: client}
}

var _ repository.Subscription = (*Subscription)(nil)

func (r *Subscription) Create(t *entity.Subscription) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "subscription" (
			"id", "ts", "created_at", "deleted_at", "status", "trial_ends_at",
			"current_period_start_at", "current_period_end_at", "canceled_at",
			"currency", "price_cents", "stripe_customer_key", "stripe_subscription_key",
			"organization_id", "plan_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Status, t.TrialEndsAt,
		t.CurrentPeriodStartAt, t.CurrentPeriodEndAt, t.CanceledAt,
		t.Currency, t.PriceCents, t.StripeCustomerKey, t.StripeSubscriptionKey,
		t.OrganizationID, t.PlanID,
	)
	return err
}

func (r *Subscription) Save(t *entity.Subscription) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "subscription" SET
			"ts" = $2, "deleted_at" = $3, "status" = $4, "trial_ends_at" = $5,
			"current_period_start_at" = $6, "current_period_end_at" = $7, "canceled_at" = $8,
			"currency" = $9, "price_cents" = $10, "stripe_customer_key" = $11, "stripe_subscription_key" = $12
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.DeletedAt, t.Status, t.TrialEndsAt,
		t.CurrentPeriodStartAt, t.CurrentPeriodEndAt, t.CanceledAt,
		t.Currency, t.PriceCents, t.StripeCustomerKey, t.StripeSubscriptionKey,
	)
	return err
}

func (r *Subscription) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "subscription" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *Subscription) FindById(id uuid.UUID) (*entity.Subscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "status", "trial_ends_at",
			"current_period_start_at", "current_period_end_at", "canceled_at",
			"currency", "price_cents", "stripe_customer_key", "stripe_subscription_key",
			"organization_id", "plan_id"
		FROM "subscription"
		WHERE "id" = $1
	`
	var t entity.Subscription
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Status, &t.TrialEndsAt,
		&t.CurrentPeriodStartAt, &t.CurrentPeriodEndAt, &t.CanceledAt,
		&t.Currency, &t.PriceCents, &t.StripeCustomerKey, &t.StripeSubscriptionKey,
		&t.OrganizationID, &t.PlanID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *Subscription) FindByOrganizationId(organizationId uuid.UUID) ([]*entity.Subscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "status", "trial_ends_at",
			"current_period_start_at", "current_period_end_at", "canceled_at",
			"currency", "price_cents", "stripe_customer_key", "stripe_subscription_key",
			"organization_id", "plan_id"
		FROM "subscription"
		WHERE "organization_id" = $1
	`
	rows, err := r.client.Query(ctx, query, organizationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.Subscription
	for rows.Next() {
		var t entity.Subscription
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Status, &t.TrialEndsAt,
			&t.CurrentPeriodStartAt, &t.CurrentPeriodEndAt, &t.CanceledAt,
			&t.Currency, &t.PriceCents, &t.StripeCustomerKey, &t.StripeSubscriptionKey,
			&t.OrganizationID, &t.PlanID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
