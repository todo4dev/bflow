// infrastructure/database/pgx/repository/artifact.go
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

type Artifact struct {
	client database.Client
}

func NewArtifact(client database.Client) *Artifact {
	return &Artifact{client: client}
}

var _ repository.Artifact = (*Artifact)(nil)

func (r *Artifact) Create(t *entity.Artifact) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "artifact" (
			"id", "ts", "created_at", "deleted_at", "kind", "name", "metadata"
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Kind, t.Name, t.Metadata,
	)
	return err
}

func (r *Artifact) Save(t *entity.Artifact) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "artifact" SET
			"ts" = $2, "deleted_at" = $3, "kind" = $4, "name" = $5, "metadata" = $6
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.DeletedAt, t.Kind, t.Name, t.Metadata,
	)
	return err
}

func (r *Artifact) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "artifact" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *Artifact) FindById(id uuid.UUID) (*entity.Artifact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "kind", "name", "metadata"
		FROM "artifact"
		WHERE "id" = $1
	`
	var t entity.Artifact
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Kind, &t.Name, &t.Metadata,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}
