// infrastructure/database/pgx/repository/organization.go
package repository

import (
	"context"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Organization struct {
	client database.Client
}

func NewOrganization(client database.Client) *Organization {
	return &Organization{client: client}
}

var _ repository.Organization = (*Organization)(nil)

func (r *Organization) Create(t *entity.Organization) error {
	query := `
		INSERT INTO "organization" (
			"id", "ts", "created_at", "deleted_at", "name", "slug", "config", "creator_account_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Name, t.Slug, t.Config, t.CreatorAccountID,
	)
	return err
}

func (r *Organization) Save(t *entity.Organization) error {
	query := `
		UPDATE "organization" SET
			"ts" = $2, "deleted_at" = $3, "name" = $4, "slug" = $5, "config" = $6
		WHERE "id" = $1
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.DeletedAt, t.Name, t.Slug, t.Config,
	)
	return err
}

func (r *Organization) Delete(id uuid.UUID) error {
	query := `DELETE FROM "organization" WHERE "id" = $1`
	_, err := r.client.Exec(context.Background(), query, id)
	return err
}

func (r *Organization) FindById(id uuid.UUID) (*entity.Organization, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "name", "slug", "config", "creator_account_id"
		FROM "organization"
		WHERE "id" = $1
	`
	var t entity.Organization
	if err := r.client.QueryRow(context.Background(), query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Name, &t.Slug, &t.Config, &t.CreatorAccountID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *Organization) FindBySlug(slug string) (*entity.Organization, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "name", "slug", "config", "creator_account_id"
		FROM "organization"
		WHERE "slug" = $1
	`
	var t entity.Organization
	if err := r.client.QueryRow(context.Background(), query, slug).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Name, &t.Slug, &t.Config, &t.CreatorAccountID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}
