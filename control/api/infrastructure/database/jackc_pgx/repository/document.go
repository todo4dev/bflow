// infrastructure/database/pgx/repository/document.go
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

type Document struct {
	client database.Client
}

func NewDocument(client database.Client) *Document {
	return &Document{client: client}
}

var _ repository.Document = (*Document)(nil)

func (r *Document) Create(t *entity.Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "document" (
			"id", "ts", "created_at", "deleted_at", "kind", "status", "title",
			"storage_key", "mimetype", "file_size", "content_sha256", "metadata",
			"replaced_document_id", "organization_id", "creator_account_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.DeletedAt, t.Kind, t.Status, t.Title,
		t.StorageKey, t.Mimetype, t.FileSize, t.ContentSHA256, t.Metadata,
		t.ReplacedDocumentID, t.OrganizationID, t.CreatorAccountID,
	)
	return err
}

func (r *Document) Save(t *entity.Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "document" SET
			"ts" = $2, "deleted_at" = $3, "status" = $4, "title" = $5,
			"storage_key" = $6, "mimetype" = $7, "file_size" = $8,
			"content_sha256" = $9, "metadata" = $10, "replaced_document_id" = $11
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.DeletedAt, t.Status, t.Title,
		t.StorageKey, t.Mimetype, t.FileSize, t.ContentSHA256, t.Metadata, t.ReplacedDocumentID,
	)
	return err
}

func (r *Document) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "document" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *Document) FindById(id uuid.UUID) (*entity.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "kind", "status", "title",
			"storage_key", "mimetype", "file_size", "content_sha256", "metadata",
			"replaced_document_id", "organization_id", "creator_account_id"
		FROM "document"
		WHERE "id" = $1
	`
	var t entity.Document
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Kind, &t.Status, &t.Title,
		&t.StorageKey, &t.Mimetype, &t.FileSize, &t.ContentSHA256, &t.Metadata,
		&t.ReplacedDocumentID, &t.OrganizationID, &t.CreatorAccountID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *Document) FindByOrganizationId(organizationId uuid.UUID) ([]*entity.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "deleted_at", "kind", "status", "title",
			"storage_key", "mimetype", "file_size", "content_sha256", "metadata",
			"replaced_document_id", "organization_id", "creator_account_id"
		FROM "document"
		WHERE "organization_id" = $1
	`
	rows, err := r.client.Query(ctx, query, organizationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.Document
	for rows.Next() {
		var t entity.Document
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.DeletedAt, &t.Kind, &t.Status, &t.Title,
			&t.StorageKey, &t.Mimetype, &t.FileSize, &t.ContentSHA256, &t.Metadata,
			&t.ReplacedDocumentID, &t.OrganizationID, &t.CreatorAccountID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
