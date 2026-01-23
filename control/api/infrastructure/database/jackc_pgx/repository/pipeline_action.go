// infrastructure/database/pgx/repository/pipeline_action.go
package repository

import (
	"context"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PipelineAction struct {
	client database.Client
}

func NewPipelineAction(client database.Client) *PipelineAction {
	return &PipelineAction{client: client}
}

var _ repository.PipelineAction = (*PipelineAction)(nil)

func (r *PipelineAction) Create(t *entity.PipelineAction) error {
	query := `
		INSERT INTO "pipeline_action" (
			"id", "ts", "created_at", "deleted_at", "kind", "status",
			"started_at", "finished_at", "error_message", "pipeline_id", "execution_cluster_agent_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Kind, t.Status,
		t.StartedAt, t.FinishedAt, t.ErrorMessage, t.PipelineID, t.ExecutionClusterAgentID,
	)
	return err
}

func (r *PipelineAction) Save(t *entity.PipelineAction) error {
	query := `
		UPDATE "pipeline_action" SET
			"ts" = $2, "deleted_at" = $3, "status" = $4, "started_at" = $5,
			"finished_at" = $6, "error_message" = $7
		WHERE "id" = $1
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.DeletedAt, t.Status, t.StartedAt, t.FinishedAt, t.ErrorMessage,
	)
	return err
}

func (r *PipelineAction) FindById(id uuid.UUID) (*entity.PipelineAction, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "kind", "status",
			"started_at", "finished_at", "error_message", "pipeline_id", "execution_cluster_agent_id"
		FROM "pipeline_action"
		WHERE "id" = $1
	`
	var t entity.PipelineAction
	if err := r.client.QueryRow(context.Background(), query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Kind, &t.Status,
		&t.StartedAt, &t.FinishedAt, &t.ErrorMessage, &t.PipelineID, &t.ExecutionClusterAgentID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *PipelineAction) FindByPipelineId(pipelineId uuid.UUID) ([]*entity.PipelineAction, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "kind", "status",
			"started_at", "finished_at", "error_message", "pipeline_id", "execution_cluster_agent_id"
		FROM "pipeline_action"
		WHERE "pipeline_id" = $1
	`
	rows, err := r.client.Query(context.Background(), query, pipelineId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.PipelineAction
	for rows.Next() {
		var t entity.PipelineAction
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Kind, &t.Status,
			&t.StartedAt, &t.FinishedAt, &t.ErrorMessage, &t.PipelineID, &t.ExecutionClusterAgentID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
