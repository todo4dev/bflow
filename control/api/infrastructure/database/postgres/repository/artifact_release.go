// infrastructure/database/pgx/repository/artifact_release.go
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

type ArtifactRelease struct {
	client database.Client
}

func NewArtifactRelease(client database.Client) *ArtifactRelease {
	return &ArtifactRelease{client: client}
}

var _ repository.ArtifactRelease = (*ArtifactRelease)(nil)

func (r *ArtifactRelease) Create(t *entity.ArtifactRelease) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "artifact_release" (
			"id", "ts", "created_at", "deleted_at", "version", "channel",
			"recommended", "notes", "published_at", "metadata", "artifact_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Version, t.Channel,
		t.Recommended, t.Notes, t.PublishedAt, t.Metadata, t.ArtifactID,
	)
	return err
}

func (r *ArtifactRelease) Save(t *entity.ArtifactRelease) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "artifact_release" SET
			"ts" = $2, "deleted_at" = $3, "version" = $4, "channel" = $5,
			"recommended" = $6, "notes" = $7, "published_at" = $8, "metadata" = $9
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.DeletedAt, t.Version, t.Channel,
		t.Recommended, t.Notes, t.PublishedAt, t.Metadata,
	)
	return err
}

func (r *ArtifactRelease) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "artifact_release" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *ArtifactRelease) FindById(id uuid.UUID) (*entity.ArtifactRelease, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "version", "channel",
			"recommended", "notes", "published_at", "metadata", "artifact_id"
		FROM "artifact_release"
		WHERE "id" = $1
	`
	var t entity.ArtifactRelease
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Version, &t.Channel,
		&t.Recommended, &t.Notes, &t.PublishedAt, &t.Metadata, &t.ArtifactID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *ArtifactRelease) FindByArtifactId(artifactId uuid.UUID) ([]*entity.ArtifactRelease, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "version", "channel",
			"recommended", "notes", "published_at", "metadata", "artifact_id"
		FROM "artifact_release"
		WHERE "artifact_id" = $1
	`
	rows, err := r.client.Query(ctx, query, artifactId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.ArtifactRelease
	for rows.Next() {
		var t entity.ArtifactRelease
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Version, &t.Channel,
			&t.Recommended, &t.Notes, &t.PublishedAt, &t.Metadata, &t.ArtifactID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
