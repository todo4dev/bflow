// infrastructure/database/pgx/repository/account.go
package repository

import (
	"context"
	"time"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
)

type Account struct {
	tableName string
	database  database.Client
}

var _ repository.Account = (*Account)(nil)

func NewAccount(database database.Client) repository.Account {
	return &Account{
		tableName: "account",
		database:  database,
	}
}

func (r *Account) Create(account *entity.Account) error {
	_, err := r.database.Exec(context.Background(), `
		INSERT INTO "`+r.tableName+`" (id, ts, created_at, deleted_at, email, email_verified_at, role)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, account.ID, account.TS, account.CreatedAt, account.DeletedAt, account.Email, account.EmailVerifiedAt, account.Role)
	return err
}

func (r *Account) Save(account *entity.Account) error {
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`"
		SET ts = $2, deleted_at = $3, email = $4, email_verified_at = $5, role = $6
		WHERE id = $1
	`, account.ID, account.TS, account.DeletedAt, account.Email, account.EmailVerifiedAt, account.Role)
	return err
}

func (r *Account) Disable(id uuid.UUID) error {
	now := time.Now()
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`" SET deleted_at = $2, ts = $2 WHERE id = $1
	`, id, now)
	return err
}

func (r *Account) FindById(id uuid.UUID) (*entity.Account, error) {
	row := r.database.QueryRow(context.Background(), `
		SELECT id, ts, created_at, deleted_at, email, email_verified_at, role
		FROM "`+r.tableName+`"
		WHERE id = $1
	`, id)

	var account entity.Account
	err := row.Scan(
		&account.ID,
		&account.TS,
		&account.CreatedAt,
		&account.DeletedAt,
		&account.Email,
		&account.EmailVerifiedAt,
		&account.Role,
	)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *Account) FindByEmail(email string) (*entity.Account, error) {
	row := r.database.QueryRow(context.Background(), `
		SELECT id, ts, created_at, deleted_at, email, email_verified_at, role
		FROM "`+r.tableName+`"
		WHERE email = $1
	`, email)

	var account entity.Account
	err := row.Scan(
		&account.ID,
		&account.TS,
		&account.CreatedAt,
		&account.DeletedAt,
		&account.Email,
		&account.EmailVerifiedAt,
		&account.Role,
	)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *Account) ExistsByEmail(email string) (bool, error) {
	row := r.database.QueryRow(context.Background(), `
		SELECT EXISTS(SELECT 1 FROM "`+r.tableName+`" WHERE email = $1)
	`, email)

	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
