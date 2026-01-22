// infrastructure/database/pgx/repository/plan.go
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

type Plan struct {
	client database.Client
}

func NewPlan(client database.Client) *Plan {
	return &Plan{client: client}
}

var _ repository.Plan = (*Plan)(nil)

func (r *Plan) Create(t *entity.Plan) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "plan" (
			"id", "ts", "created_at", "deleted_at", "code", "name",
			"interval", "currency", "price_cents", "config"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Code, t.Name,
		t.Interval, t.Currency, t.PriceCents, t.Config,
	)
	return err
}

func (r *Plan) Save(t *entity.Plan) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "plan" SET
			"ts" = $2, "deleted_at" = $3, "code" = $4, "name" = $5,
			"interval" = $6, "currency" = $7, "price_cents" = $8, "config" = $9
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.DeletedAt, t.Code, t.Name,
		t.Interval, t.Currency, t.PriceCents, t.Config,
	)
	return err
}

func (r *Plan) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "plan" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *Plan) FindById(id uuid.UUID) (*entity.Plan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "code", "name",
			"interval", "currency", "price_cents", "config"
		FROM "plan"
		WHERE "id" = $1
	`
	var t entity.Plan
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Code, &t.Name,
		&t.Interval, &t.Currency, &t.PriceCents, &t.Config,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *Plan) FindByCode(code string) (*entity.Plan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "code", "name",
			"interval", "currency", "price_cents", "config"
		FROM "plan"
		WHERE "code" = $1
	`
	var t entity.Plan
	if err := r.client.QueryRow(ctx, query, code).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Code, &t.Name,
		&t.Interval, &t.Currency, &t.PriceCents, &t.Config,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}
