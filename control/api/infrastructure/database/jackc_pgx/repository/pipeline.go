// infrastructure/database/pgx/repository/pipeline.go
package repository

import (
	"context"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Pipeline struct {
	client database.Client
}

func NewPipeline(client database.Client) *Pipeline {
	return &Pipeline{client: client}
}

var _ repository.Pipeline = (*Pipeline)(nil)

func (r *Pipeline) Create(t *entity.Pipeline) error {
	query := `
		INSERT INTO "pipeline" (
			"id", "ts", "created_at", "deleted_at", "kind", "status", "payload",
			"started_at", "finished_at", "error_message", "target_artifact_release_id",
			"cluster_runtime_id", "requester_account_id", "organization_id", "previous_pipeline_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Kind, t.Status, t.Payload,
		t.StartedAt, t.FinishedAt, t.ErrorMessage, t.TargetArtifactReleaseID,
		t.ClusterRuntimeID, t.RequesterAccountID, t.OrganizationID, t.PreviousPipelineID,
	)
	return err
}

func (r *Pipeline) Save(t *entity.Pipeline) error {
	query := `
		UPDATE "pipeline" SET
			"ts" = $2, "deleted_at" = $3, "status" = $4, "started_at" = $5,
			"finished_at" = $6, "error_message" = $7
		WHERE "id" = $1
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.DeletedAt, t.Status, t.StartedAt, t.FinishedAt, t.ErrorMessage,
	)
	return err
}

func (r *Pipeline) FindById(id uuid.UUID) (*entity.Pipeline, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "kind", "status", "payload",
			"started_at", "finished_at", "error_message", "target_artifact_release_id",
			"cluster_runtime_id", "requester_account_id", "organization_id", "previous_pipeline_id"
		FROM "pipeline"
		WHERE "id" = $1
	`
	var t entity.Pipeline
	if err := r.client.QueryRow(context.Background(), query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Kind, &t.Status, &t.Payload,
		&t.StartedAt, &t.FinishedAt, &t.ErrorMessage, &t.TargetArtifactReleaseID,
		&t.ClusterRuntimeID, &t.RequesterAccountID, &t.OrganizationID, &t.PreviousPipelineID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *Pipeline) FindByOrganizationId(organizationId uuid.UUID) ([]*entity.Pipeline, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "kind", "status", "payload",
			"started_at", "finished_at", "error_message", "target_artifact_release_id",
			"cluster_runtime_id", "requester_account_id", "organization_id", "previous_pipeline_id"
		FROM "pipeline"
		WHERE "organization_id" = $1
	`
	rows, err := r.client.Query(context.Background(), query, organizationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.Pipeline
	for rows.Next() {
		var t entity.Pipeline
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Kind, &t.Status, &t.Payload,
			&t.StartedAt, &t.FinishedAt, &t.ErrorMessage, &t.TargetArtifactReleaseID,
			&t.ClusterRuntimeID, &t.RequesterAccountID, &t.OrganizationID, &t.PreviousPipelineID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}

func (r *Pipeline) FindByClusterRuntimeId(runtimeId uuid.UUID) ([]*entity.Pipeline, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "kind", "status", "payload",
			"started_at", "finished_at", "error_message", "target_artifact_release_id",
			"cluster_runtime_id", "requester_account_id", "organization_id", "previous_pipeline_id"
		FROM "pipeline"
		WHERE "cluster_runtime_id" = $1
	`
	rows, err := r.client.Query(context.Background(), query, runtimeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.Pipeline
	for rows.Next() {
		var t entity.Pipeline
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Kind, &t.Status, &t.Payload,
			&t.StartedAt, &t.FinishedAt, &t.ErrorMessage, &t.TargetArtifactReleaseID,
			&t.ClusterRuntimeID, &t.RequesterAccountID, &t.OrganizationID, &t.PreviousPipelineID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
