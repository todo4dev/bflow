// infrastructure/database/pgx/repository/account_activity.go
package repository

import (
	"context"

	"src/domain/entity"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
)

type AccountActivity struct {
	tableName string
	database  database.Client
}

var _ repository.AccountActivity = (*AccountActivity)(nil)

func NewAccountActivity(database database.Client) repository.AccountActivity {
	return &AccountActivity{
		tableName: "account_activity",
		database:  database,
	}
}

func (r *AccountActivity) Create(activity *entity.AccountActivity) error {
	query := `
		INSERT INTO "` + r.tableName + `" (
			"id", "ts", "created_at", "kind", "message", "metadata", "account_id", 
			"organization_id", "cluster_id"
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)
	`
	_, err := r.database.Exec(context.Background(), query,
		activity.ID, activity.TS, activity.CreatedAt, activity.Kind,
		activity.Message, activity.Metadata, activity.AccountID,
		activity.OrganizationID, activity.ClusterID,
	)
	return err
}

func (r *AccountActivity) FindById(id uuid.UUID) (*entity.AccountActivity, error) {
	query := `
		SELECT "id", "ts", "created_at", "kind", "message", "metadata", "account_id", 
		"organization_id", "cluster_id"
		FROM "` + r.tableName + `"
		WHERE "id" = $1
	`
	row := r.database.QueryRow(context.Background(), query, id)

	var activity entity.AccountActivity
	err := row.Scan(
		&activity.ID, &activity.TS, &activity.CreatedAt, &activity.Kind,
		&activity.Message, &activity.Metadata, &activity.AccountID,
		&activity.OrganizationID, &activity.ClusterID,
	)
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *AccountActivity) FindByAccountId(accountId uuid.UUID) ([]*entity.AccountActivity, error) {
	query := `
		SELECT "id", "ts", "created_at", "kind", "message", "metadata", "account_id", 
		"organization_id", "cluster_id"
		FROM "` + r.tableName + `"
		WHERE "account_id" = $1
	`
	rows, err := r.database.Query(context.Background(), query, accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []*entity.AccountActivity
	for rows.Next() {
		var activity entity.AccountActivity
		err := rows.Scan(
			&activity.ID, &activity.TS, &activity.CreatedAt, &activity.Kind,
			&activity.Message, &activity.Metadata, &activity.AccountID,
			&activity.OrganizationID, &activity.ClusterID,
		)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &activity)
	}
	return activities, nil
}

func (r *AccountActivity) FindByClusterId(clusterId uuid.UUID) ([]*entity.AccountActivity, error) {
	query := `
		SELECT "id", "ts", "created_at", "kind", "message", "metadata", "account_id", 
		"organization_id", "cluster_id"
		FROM "` + r.tableName + `"
		WHERE "cluster_id" = $1
	`
	rows, err := r.database.Query(context.Background(), query, clusterId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []*entity.AccountActivity
	for rows.Next() {
		var activity entity.AccountActivity
		err := rows.Scan(
			&activity.ID, &activity.TS, &activity.CreatedAt, &activity.Kind,
			&activity.Message, &activity.Metadata, &activity.AccountID,
			&activity.OrganizationID, &activity.ClusterID,
		)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &activity)
	}
	return activities, nil
}
