// infrastructure/database/pgx/repository/account_certificate.go
package repository

import (
	"context"
	"time"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
)

type AccountCertificate struct {
	tableName string
	database  database.Client
}

var _ repository.AccountCertificate = (*AccountCertificate)(nil)

func NewAccountCertificate(database database.Client) repository.AccountCertificate {
	return &AccountCertificate{
		tableName: "account_certificate",
		database:  database,
	}
}

func (r *AccountCertificate) Create(cert *entity.AccountCertificate) error {
	_, err := r.database.Exec(context.Background(), `
		INSERT INTO "`+r.tableName+`" (id, ts, created_at, deleted_at, expires_at, display_name, document_number, owner_name, thumbprint, is_active, account_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`, cert.ID, cert.TS, cert.CreatedAt, cert.DeletedAt, cert.ExpiresAt, cert.DisplayName, cert.DocumentNumber, cert.OwnerName, cert.Thumbprint, cert.IsActive, cert.AccountID)
	return err
}

func (r *AccountCertificate) Save(cert *entity.AccountCertificate) error {
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`"
		SET ts = $2, deleted_at = $3, expires_at = $4, display_name = $5, document_number = $6, owner_name = $7, thumbprint = $8, is_active = $9
		WHERE id = $1
	`, cert.ID, cert.TS, cert.DeletedAt, cert.ExpiresAt, cert.DisplayName, cert.DocumentNumber, cert.OwnerName, cert.Thumbprint, cert.IsActive)
	return err
}

func (r *AccountCertificate) Delete(id uuid.UUID) error {
	now := time.Now()
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`" SET deleted_at = $2, ts = $2, is_active = false WHERE id = $1
	`, id, now)
	return err
}

func (r *AccountCertificate) FindById(id uuid.UUID) (*entity.AccountCertificate, error) {
	row := r.database.QueryRow(context.Background(), `
		SELECT id, ts, created_at, deleted_at, expires_at, display_name, document_number, owner_name, thumbprint, is_active, account_id
		FROM "`+r.tableName+`"
		WHERE id = $1
	`, id)

	var cert entity.AccountCertificate
	err := row.Scan(
		&cert.ID,
		&cert.TS,
		&cert.CreatedAt,
		&cert.DeletedAt,
		&cert.ExpiresAt,
		&cert.DisplayName,
		&cert.DocumentNumber,
		&cert.OwnerName,
		&cert.Thumbprint,
		&cert.IsActive,
		&cert.AccountID,
	)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func (r *AccountCertificate) FindByAccountId(accountId uuid.UUID) ([]*entity.AccountCertificate, error) {
	rows, err := r.database.Query(context.Background(), `
		SELECT id, ts, created_at, deleted_at, expires_at, display_name, document_number, owner_name, thumbprint, is_active, account_id
		FROM "`+r.tableName+`"
		WHERE account_id = $1
	`, accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var certs []*entity.AccountCertificate
	for rows.Next() {
		var cert entity.AccountCertificate
		err := rows.Scan(
			&cert.ID,
			&cert.TS,
			&cert.CreatedAt,
			&cert.DeletedAt,
			&cert.ExpiresAt,
			&cert.DisplayName,
			&cert.DocumentNumber,
			&cert.OwnerName,
			&cert.Thumbprint,
			&cert.IsActive,
			&cert.AccountID,
		)
		if err != nil {
			return nil, err
		}
		certs = append(certs, &cert)
	}
	return certs, nil
}

func (r *AccountCertificate) FindActiveByAccountId(accountId uuid.UUID) (*entity.AccountCertificate, error) {
	row := r.database.QueryRow(context.Background(), `
		SELECT id, ts, created_at, deleted_at, expires_at, display_name, document_number, owner_name, thumbprint, is_active, account_id
		FROM "`+r.tableName+`"
		WHERE account_id = $1 AND is_active = true AND deleted_at IS NULL
	`, accountId)

	var cert entity.AccountCertificate
	err := row.Scan(
		&cert.ID,
		&cert.TS,
		&cert.CreatedAt,
		&cert.DeletedAt,
		&cert.ExpiresAt,
		&cert.DisplayName,
		&cert.DocumentNumber,
		&cert.OwnerName,
		&cert.Thumbprint,
		&cert.IsActive,
		&cert.AccountID,
	)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}
