// infrastructure/database/pgx/repository/account_credential.go
package repository

import (
	"context"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
)

type AccountCredential struct {
	tableName string
	database  database.Client
}

var _ repository.AccountCredential = (*AccountCredential)(nil)

func NewAccountCredential(database database.Client) repository.AccountCredential {
	return &AccountCredential{
		tableName: "account_credential",
		database:  database,
	}
}

func (r *AccountCredential) Create(credential *entity.AccountCredential) error {
	_, err := r.database.Exec(context.Background(), `
		INSERT INTO "`+r.tableName+`" (id, ts, created_at, password_hash, account_id)
		VALUES ($1, $2, $3, $4, $5)
	`, credential.ID, credential.TS, credential.CreatedAt, credential.PasswordHash, credential.AccountID)
	return err
}

func (r *AccountCredential) Save(credential *entity.AccountCredential) error {
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`"
		SET ts = $2, password_hash = $3
		WHERE id = $1
	`, credential.ID, credential.TS, credential.PasswordHash)
	return err
}

func (r *AccountCredential) FindByAccountId(accountId uuid.UUID) (*entity.AccountCredential, error) {
	row := r.database.QueryRow(context.Background(), `
		SELECT id, ts, created_at, password_hash, account_id
		FROM "`+r.tableName+`"
		WHERE account_id = $1
	`, accountId)

	var credential entity.AccountCredential
	err := row.Scan(
		&credential.ID,
		&credential.TS,
		&credential.CreatedAt,
		&credential.PasswordHash,
		&credential.AccountID,
	)
	if err != nil {
		return nil, err
	}
	return &credential, nil
}
