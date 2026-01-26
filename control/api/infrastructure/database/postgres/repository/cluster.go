// infrastructure/database/pgx/repository/cluster.go
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

type Cluster struct {
	client database.Client
}

func NewCluster(client database.Client) *Cluster {
	return &Cluster{client: client}
}

var _ repository.Cluster = (*Cluster)(nil)

func (r *Cluster) Create(t *entity.Cluster) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "cluster" (
			"id", "ts", "created_at", "deleted_at", "state", "promoted_at", "legacy_at",
			"name", "namespace", "provider", "external_id", "kubernetes_version", "metadata", "organization_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.State, t.PromotedAt, t.LegacyAt,
		t.Name, t.Namespace, t.Provider, t.ExternalID, t.KubernetesVersion, t.Metadata, t.OrganizationID,
	)
	return err
}

func (r *Cluster) Save(t *entity.Cluster) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "cluster" SET
			"ts" = $2, "deleted_at" = $3, "state" = $4, "promoted_at" = $5, "legacy_at" = $6,
			"name" = $7, "namespace" = $8, "provider" = $9, "external_id" = $10,
			"kubernetes_version" = $11, "metadata" = $12
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.DeletedAt, t.State, t.PromotedAt, t.LegacyAt,
		t.Name, t.Namespace, t.Provider, t.ExternalID, t.KubernetesVersion, t.Metadata,
	)
	return err
}

func (r *Cluster) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "cluster" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *Cluster) FindById(id uuid.UUID) (*entity.Cluster, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "state", "promoted_at", "legacy_at",
			"name", "namespace", "provider", "external_id", "kubernetes_version", "metadata", "organization_id"
		FROM "cluster"
		WHERE "id" = $1
	`
	var t entity.Cluster
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.State, &t.PromotedAt, &t.LegacyAt,
		&t.Name, &t.Namespace, &t.Provider, &t.ExternalID, &t.KubernetesVersion, &t.Metadata, &t.OrganizationID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *Cluster) FindByOrganizationId(organizationId uuid.UUID) ([]*entity.Cluster, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "state", "promoted_at", "legacy_at",
			"name", "namespace", "provider", "external_id", "kubernetes_version", "metadata", "organization_id"
		FROM "cluster"
		WHERE "organization_id" = $1
	`
	rows, err := r.client.Query(ctx, query, organizationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.Cluster
	for rows.Next() {
		var t entity.Cluster
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.State, &t.PromotedAt, &t.LegacyAt,
			&t.Name, &t.Namespace, &t.Provider, &t.ExternalID, &t.KubernetesVersion, &t.Metadata, &t.OrganizationID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
