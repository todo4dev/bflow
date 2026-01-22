// infrastructure/database/pgx/repository/organization_invite.go
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

type OrganizationInvite struct {
	client database.Client
}

func NewOrganizationInvite(client database.Client) *OrganizationInvite {
	return &OrganizationInvite{client: client}
}

var _ repository.OrganizationInvite = (*OrganizationInvite)(nil)

func (r *OrganizationInvite) Create(t *entity.OrganizationInvite) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "organization_invite" (
			"id", "ts", "created_at", "code", "email", "role", "status",
			"expires_at", "accepted_at", "organization_id", "creator_account_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.Code, t.Email, t.Role, t.Status,
		t.ExpiresAt, t.AcceptedAt, t.OrganizationID, t.CreatorAccountID,
	)
	return err
}

func (r *OrganizationInvite) Save(t *entity.OrganizationInvite) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "organization_invite" SET
			"ts" = $2, "role" = $3, "status" = $4, "expires_at" = $5, "accepted_at" = $6
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.Role, t.Status, t.ExpiresAt, t.AcceptedAt,
	)
	return err
}

func (r *OrganizationInvite) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "organization_invite" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *OrganizationInvite) FindById(id uuid.UUID) (*entity.OrganizationInvite, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "code", "email", "role", "status",
			"expires_at", "accepted_at", "organization_id", "creator_account_id"
		FROM "organization_invite"
		WHERE "id" = $1
	`
	var t entity.OrganizationInvite
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.Code, &t.Email, &t.Role, &t.Status,
		&t.ExpiresAt, &t.AcceptedAt, &t.OrganizationID, &t.CreatorAccountID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *OrganizationInvite) FindByCode(code string) (*entity.OrganizationInvite, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "code", "email", "role", "status",
			"expires_at", "accepted_at", "organization_id", "creator_account_id"
		FROM "organization_invite"
		WHERE "code" = $1
	`
	var t entity.OrganizationInvite
	if err := r.client.QueryRow(ctx, query, code).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.Code, &t.Email, &t.Role, &t.Status,
		&t.ExpiresAt, &t.AcceptedAt, &t.OrganizationID, &t.CreatorAccountID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *OrganizationInvite) FindByOrganizationId(organizationId uuid.UUID) ([]*entity.OrganizationInvite, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "code", "email", "role", "status",
			"expires_at", "accepted_at", "organization_id", "creator_account_id"
		FROM "organization_invite"
		WHERE "organization_id" = $1
	`
	rows, err := r.client.Query(ctx, query, organizationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.OrganizationInvite
	for rows.Next() {
		var t entity.OrganizationInvite
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.Code, &t.Email, &t.Role, &t.Status,
			&t.ExpiresAt, &t.AcceptedAt, &t.OrganizationID, &t.CreatorAccountID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
