// infrastructure/database/pgx/repository/cluster_agent.go
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

type ClusterAgent struct {
	client database.Client
}

func NewClusterAgent(client database.Client) *ClusterAgent {
	return &ClusterAgent{client: client}
}

var _ repository.ClusterAgent = (*ClusterAgent)(nil)

func (r *ClusterAgent) Create(t *entity.ClusterAgent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "cluster_agent" (
			"id", "ts", "created_at", "deleted_at", "status", "version", "last_seen_at",
			"public_key", "metadata", "cluster_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Status, t.Version, t.LastSeenAt,
		t.PublicKey, t.Metadata, t.ClusterID,
	)
	return err
}

func (r *ClusterAgent) Save(t *entity.ClusterAgent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "cluster_agent" SET
			"ts" = $2, "deleted_at" = $3, "status" = $4, "version" = $5,
			"last_seen_at" = $6, "public_key" = $7, "metadata" = $8
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.DeletedAt, t.Status, t.Version, t.LastSeenAt, t.PublicKey, t.Metadata,
	)
	return err
}

func (r *ClusterAgent) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "cluster_agent" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *ClusterAgent) FindById(id uuid.UUID) (*entity.ClusterAgent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "status", "version", "last_seen_at",
			"public_key", "metadata", "cluster_id"
		FROM "cluster_agent"
		WHERE "id" = $1
	`
	var t entity.ClusterAgent
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Status, &t.Version, &t.LastSeenAt,
		&t.PublicKey, &t.Metadata, &t.ClusterID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *ClusterAgent) FindByClusterId(clusterId uuid.UUID) ([]*entity.ClusterAgent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "status", "version", "last_seen_at",
			"public_key", "metadata", "cluster_id"
		FROM "cluster_agent"
		WHERE "cluster_id" = $1
	`
	rows, err := r.client.Query(ctx, query, clusterId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.ClusterAgent
	for rows.Next() {
		var t entity.ClusterAgent
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Status, &t.Version, &t.LastSeenAt,
			&t.PublicKey, &t.Metadata, &t.ClusterID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
