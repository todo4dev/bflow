package identity

import (
	"context"
	"time"

	"src/domain/identity/entity"
	"src/domain/identity/repository"
	"src/port/database"

	"github.com/google/uuid"
)

type AccountRepository struct {
	tableName string
	database  database.Client
}

var _ repository.Account = (*AccountRepository)(nil)

func NewAccountRepository(database database.Client) repository.Account {
	return &AccountRepository{
		tableName: "account",
		database:  database,
	}
}

func (r *AccountRepository) Create(account *entity.Account) error {
	_, err := r.database.Exec(context.Background(), `
		INSERT INTO "`+r.tableName+`" (id, ts, created_at, deleted_at, email, email_verified_at, role)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, account.ID, account.TS, account.CreatedAt, account.DeletedAt, account.Email, account.EmailVerifiedAt, account.Role)
	return err
}

func (r *AccountRepository) Save(account *entity.Account) error {
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`"
		SET ts = $2, deleted_at = $3, email = $4, email_verified_at = $5, role = $6
		WHERE id = $1
	`, account.ID, account.TS, account.DeletedAt, account.Email, account.EmailVerifiedAt, account.Role)
	return err
}

func (r *AccountRepository) Disable(id uuid.UUID) error {
	now := time.Now()
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`" SET deleted_at = $2, ts = $2 WHERE id = $1
	`, id, now)
	return err
}

func (r *AccountRepository) FindById(id uuid.UUID) (*entity.Account, error) {
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

func (r *AccountRepository) FindByEmail(email string) (*entity.Account, error) {
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

func (r *AccountRepository) ExistsByEmail(email string) (bool, error) {
	row := r.database.QueryRow(context.Background(), `
		SELECT EXISTS(SELECT 1 FROM "`+r.tableName+`" WHERE email = $1)
	`, email)

	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
