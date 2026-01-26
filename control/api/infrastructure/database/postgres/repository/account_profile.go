// infrastructure/database/pgx/repository/account_profile.go
package repository

import (
	"context"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
)

type AccountProfile struct {
	tableName string
	database  database.Client
}

var _ repository.AccountProfile = (*AccountProfile)(nil)

func NewAccountProfile(database database.Client) repository.AccountProfile {
	return &AccountProfile{
		tableName: "account_profile",
		database:  database,
	}
}

func (r *AccountProfile) Create(profile *entity.AccountProfile) error {
	_, err := r.database.Exec(context.Background(), `
		INSERT INTO "`+r.tableName+`" (id, ts, given_name, family_name, language, timezone, account_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, profile.ID, profile.TS, profile.GivenName, profile.FamilyName, profile.Language, profile.Timezone, profile.AccountID)
	return err
}

func (r *AccountProfile) Save(profile *entity.AccountProfile) error {
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`"
		SET ts = $2, given_name = $3, family_name = $4, language = $5, timezone = $6
		WHERE id = $1
	`, profile.ID, profile.TS, profile.GivenName, profile.FamilyName, profile.Language, profile.Timezone)
	return err
}

func (r *AccountProfile) FindByAccountId(accountId uuid.UUID) (*entity.AccountProfile, error) {
	row := r.database.QueryRow(context.Background(), `
		SELECT id, ts, given_name, family_name, language, timezone, account_id
		FROM "`+r.tableName+`"
		WHERE account_id = $1
	`, accountId)

	var profile entity.AccountProfile
	err := row.Scan(
		&profile.ID,
		&profile.TS,
		&profile.GivenName,
		&profile.FamilyName,
		&profile.Language,
		&profile.Timezone,
		&profile.AccountID,
	)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
