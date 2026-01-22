// infrastructure/database/pgx/repository/document_signature.go
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

type DocumentSignature struct {
	client database.Client
}

func NewDocumentSignature(client database.Client) *DocumentSignature {
	return &DocumentSignature{client: client}
}

var _ repository.DocumentSignature = (*DocumentSignature)(nil)

func (r *DocumentSignature) Create(t *entity.DocumentSignature) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "document_signature" (
			"id", "ts", "created_at", "state", "signed_at", "failure_reason",
			"value", "hash", "metadata", "document_id", "account_certificate_id", "account_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.State, t.SignedAt, t.FailureReason,
		t.Value, t.Hash, t.Metadata, t.DocumentID, t.AccountCertificateID, t.AccountID,
	)
	return err
}

func (r *DocumentSignature) Save(t *entity.DocumentSignature) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "document_signature" SET
			"ts" = $2, "state" = $3, "signed_at" = $4, "failure_reason" = $5,
			"value" = $6, "hash" = $7, "metadata" = $8
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.State, t.SignedAt, t.FailureReason, t.Value, t.Hash, t.Metadata,
	)
	return err
}

func (r *DocumentSignature) FindById(id uuid.UUID) (*entity.DocumentSignature, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "state", "signed_at", "failure_reason",
			"value", "hash", "metadata", "document_id", "account_certificate_id", "account_id"
		FROM "document_signature"
		WHERE "id" = $1
	`
	var t entity.DocumentSignature
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.State, &t.SignedAt, &t.FailureReason,
		&t.Value, &t.Hash, &t.Metadata, &t.DocumentID, &t.AccountCertificateID, &t.AccountID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *DocumentSignature) FindByDocumentId(documentId uuid.UUID) ([]*entity.DocumentSignature, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "state", "signed_at", "failure_reason",
			"value", "hash", "metadata", "document_id", "account_certificate_id", "account_id"
		FROM "document_signature"
		WHERE "document_id" = $1
	`
	rows, err := r.client.Query(ctx, query, documentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.DocumentSignature
	for rows.Next() {
		var t entity.DocumentSignature
		if err := rows.Scan(
			&t.ID, &t.TS, &t.CreatedAt, &t.State, &t.SignedAt, &t.FailureReason,
			&t.Value, &t.Hash, &t.Metadata, &t.DocumentID, &t.AccountCertificateID, &t.AccountID,
		); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}
