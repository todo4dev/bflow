// infrastructure/database/pgx/repository/account_preference.go
package repository

import (
	"context"
	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
)

type AccountPreference struct {
	tableName string
	client    database.Client
}

func NewAccountPreference(client database.Client) repository.AccountPreference {
	return &AccountPreference{
		tableName: "account_preference",
		client:    client,
	}
}

func (r *AccountPreference) Create(preference *entity.AccountPreference) error {
	query := `
		INSERT INTO "` + r.tableName + `" (
			"id", "ts", "theme", "notify_on_pipeline_success", 
			"notify_on_infra_alerts", "account_id"
		) VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.client.Exec(context.Background(), query,
		preference.ID,
		preference.TS,
		preference.Theme,
		preference.NotifyOnPipelineSuccess,
		preference.NotifyOnInfraAlerts,
		preference.AccountID,
	)

	return err
}

func (r *AccountPreference) Save(preference *entity.AccountPreference) error {
	query := `
		UPDATE "` + r.tableName + `" SET
			"ts" = $1,
			"theme" = $2,
			"notify_on_pipeline_success" = $3,
			"notify_on_infra_alerts" = $4
		WHERE "id" = $5
	`
	_, err := r.client.Exec(context.Background(), query,
		preference.TS, preference.Theme, preference.NotifyOnPipelineSuccess,
		preference.NotifyOnInfraAlerts, preference.ID,
	)

	return err
}

func (r *AccountPreference) FindByAccountId(accountId uuid.UUID) (*entity.AccountPreference, error) {
	query := `
		SELECT 
			"id", "ts", "theme", "notify_on_pipeline_success", "notify_on_infra_alerts", 
			"account_id"
		FROM "` + r.tableName + `"
		WHERE "account_id" = $1
	`

	var preference entity.AccountPreference
	err := r.client.QueryRow(context.Background(), query, accountId).Scan(
		&preference.ID, &preference.TS, &preference.Theme,
		&preference.NotifyOnPipelineSuccess, &preference.NotifyOnInfraAlerts,
		&preference.AccountID,
	)
	if err != nil {
		return nil, err
	}
	return &preference, nil
}
