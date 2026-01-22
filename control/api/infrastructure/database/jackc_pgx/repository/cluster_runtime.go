// infrastructure/database/pgx/repository/cluster_runtime.go
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

type ClusterRuntime struct {
	client database.Client
}

func NewClusterRuntime(client database.Client) *ClusterRuntime {
	return &ClusterRuntime{client: client}
}

var _ repository.ClusterRuntime = (*ClusterRuntime)(nil)

func (r *ClusterRuntime) Create(t *entity.ClusterRuntime) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "cluster_runtime" (
			"id", "ts", "created_at", "deleted_at", "state", "readonly", "last_deployed_at",
			"config", "current_artifact_release_id", "cluster_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.State, t.ReadOnly, t.LastDeployedAt,
		t.Config, t.CurrentArtifactReleaseID, t.ClusterID,
	)
	return err
}

func (r *ClusterRuntime) Save(t *entity.ClusterRuntime) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "cluster_runtime" SET
			"ts" = $2, "deleted_at" = $3, "state" = $4, "readonly" = $5, "last_deployed_at" = $6,
			"config" = $7, "current_artifact_release_id" = $8
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.DeletedAt, t.State, t.ReadOnly, t.LastDeployedAt,
		t.Config, t.CurrentArtifactReleaseID,
	)
	return err
}

func (r *ClusterRuntime) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "cluster_runtime" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *ClusterRuntime) FindById(id uuid.UUID) (*entity.ClusterRuntime, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "state", "readonly", "last_deployed_at",
			"config", "current_artifact_release_id", "cluster_id"
		FROM "cluster_runtime"
		WHERE "id" = $1
	`
	var t entity.ClusterRuntime
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.State, &t.ReadOnly, &t.LastDeployedAt,
		&t.Config, &t.CurrentArtifactReleaseID, &t.ClusterID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *ClusterRuntime) FindByClusterId(clusterId uuid.UUID) ([]*entity.ClusterRuntime, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "state", "readonly", "last_deployed_at",
			"config", "current_artifact_release_id", "cluster_id"
		FROM "cluster_runtime"
		WHERE "cluster_id" = $1
	`
	rows, err := r.client.Query(ctx, query, clusterId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.ClusterRuntime
	for rows.Next() {
		var t entity.ClusterRuntime
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.State, &t.ReadOnly, &t.LastDeployedAt,
			&t.Config, &t.CurrentArtifactReleaseID, &t.ClusterID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
