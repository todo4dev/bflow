// infrastructure/database/pgx/repository/pipeline_action_stage.go
package repository

import (
	"context"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PipelineActionStage struct {
	client database.Client
}

func NewPipelineActionStage(client database.Client) *PipelineActionStage {
	return &PipelineActionStage{client: client}
}

var _ repository.PipelineActionStage = (*PipelineActionStage)(nil)

func (r *PipelineActionStage) Create(t *entity.PipelineActionStage) error {
	query := `
		INSERT INTO "pipeline_action_stage" (
			"id", "ts", "created_at", "deleted_at", "name", "position", "status",
			"started_at", "finished_at", "summary", "output_meta", "pipeline_action_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Name, t.Position, t.Status,
		t.StartedAt, t.FinishedAt, t.Summary, t.OutputMeta, t.PipelineActionID,
	)
	return err
}

func (r *PipelineActionStage) Save(t *entity.PipelineActionStage) error {
	query := `
		UPDATE "pipeline_action_stage" SET
			"ts" = $2, "deleted_at" = $3, "status" = $4, "started_at" = $5,
			"finished_at" = $6, "summary" = $7, "output_meta" = $8
		WHERE "id" = $1
	`
	_, err := r.client.Exec(context.Background(), query,
		t.ID, t.TS, t.DeletedAt, t.Status, t.StartedAt, t.FinishedAt, t.Summary, t.OutputMeta,
	)
	return err
}

func (r *PipelineActionStage) FindById(id uuid.UUID) (*entity.PipelineActionStage, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "name", "position", "status",
			"started_at", "finished_at", "summary", "output_meta", "pipeline_action_id"
		FROM "pipeline_action_stage"
		WHERE "id" = $1
	`
	var t entity.PipelineActionStage
	if err := r.client.QueryRow(context.Background(), query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Name, &t.Position, &t.Status,
		&t.StartedAt, &t.FinishedAt, &t.Summary, &t.OutputMeta, &t.PipelineActionID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *PipelineActionStage) FindByPipelineActionId(actionId uuid.UUID) ([]*entity.PipelineActionStage, error) {
	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "name", "position", "status",
			"started_at", "finished_at", "summary", "output_meta", "pipeline_action_id"
		FROM "pipeline_action_stage"
		WHERE "pipeline_action_id" = $1
	`
	rows, err := r.client.Query(context.Background(), query, actionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.PipelineActionStage
	for rows.Next() {
		var t entity.PipelineActionStage
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Name, &t.Position, &t.Status,
			&t.StartedAt, &t.FinishedAt, &t.Summary, &t.OutputMeta, &t.PipelineActionID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
