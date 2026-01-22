// infrastructure/database/pgx/repository/organization_membership.go
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

type OrganizationMembership struct {
	client database.Client
}

func NewOrganizationMembership(client database.Client) *OrganizationMembership {
	return &OrganizationMembership{client: client}
}

var _ repository.OrganizationMembership = (*OrganizationMembership)(nil)

func (r *OrganizationMembership) Create(t *entity.OrganizationMembership) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "organization_membership" (
			"id", "ts", "created_at", "role", "account_id", "organization_id"
		) VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.Role, t.AccountID, t.OrganizationID,
	)
	return err
}

func (r *OrganizationMembership) Save(t *entity.OrganizationMembership) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "organization_membership" SET
			"ts" = $2, "role" = $3
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.Role,
	)
	return err
}

func (r *OrganizationMembership) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "organization_membership" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *OrganizationMembership) FindById(id uuid.UUID) (*entity.OrganizationMembership, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "role", "account_id", "organization_id"
		FROM "organization_membership"
		WHERE "id" = $1
	`
	var t entity.OrganizationMembership
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.Role, &t.AccountID, &t.OrganizationID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *OrganizationMembership) FindByAccountId(accountId uuid.UUID) ([]*entity.OrganizationMembership, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "role", "account_id", "organization_id"
		FROM "organization_membership"
		WHERE "account_id" = $1
	`
	rows, err := r.client.Query(ctx, query, accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.OrganizationMembership
	for rows.Next() {
		var t entity.OrganizationMembership
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.Role, &t.AccountID, &t.OrganizationID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}

func (r *OrganizationMembership) FindByOrganizationId(organizationId uuid.UUID) ([]*entity.OrganizationMembership, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "role", "account_id", "organization_id"
		FROM "organization_membership"
		WHERE "organization_id" = $1
	`
	rows, err := r.client.Query(ctx, query, organizationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.OrganizationMembership
	for rows.Next() {
		var t entity.OrganizationMembership
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.Role, &t.AccountID, &t.OrganizationID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
