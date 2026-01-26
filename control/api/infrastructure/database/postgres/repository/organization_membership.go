// infrastructure/database/pgx/repository/organization_membership.go
package repository

import (
	"context"

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

	query := `
		INSERT INTO "organization_membership" (
			"id", "ts", "created_at", "role", "account_id", "organization_id"
		) VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.CreatedAt, t.Role, t.AccountID, t.OrganizationID,
	)
	return err
}

func (r *OrganizationMembership) Save(t *entity.OrganizationMembership) error {
	query := `
		UPDATE "organization_membership" SET
			"ts" = $2, "role" = $3
		WHERE "id" = $1
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.Role,
	)
	return err
}

func (r *OrganizationMembership) Delete(id uuid.UUID) error {
	query := `DELETE FROM "organization_membership" WHERE "id" = $1`
	_, err := r.client.Exec(context.Background(), query, id)
	return err
}

func (r *OrganizationMembership) FindById(id uuid.UUID) (*entity.OrganizationMembership, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "role", "account_id", "organization_id"
		FROM "organization_membership"
		WHERE "id" = $1
	`
	var t entity.OrganizationMembership
	if err := r.client.QueryRow(context.Background(), query, id).Scan(
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
	query := `
		SELECT
			"id", "ts", "created_at", "role", "account_id", "organization_id"
		FROM "organization_membership"
		WHERE "account_id" = $1
	`
	rows, err := r.client.Query(context.Background(), query, accountId)
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
	query := `
		SELECT
			"id", "ts", "created_at", "role", "account_id", "organization_id"
		FROM "organization_membership"
		WHERE "organization_id" = $1
	`
	rows, err := r.client.Query(context.Background(), query, organizationId)
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
