// infrastructure/database/pgx/repository/account_notification.go
package repository

import (
	"context"
	"time"

	"src/domain/entity"
	"src/domain/enum"
	"src/domain/repository"
	"src/port/database"

	"github.com/google/uuid"
)

type AccountNotification struct {
	tableName string
	database  database.Client
}

var _ repository.AccountNotification = (*AccountNotification)(nil)

func NewAccountNotification(database database.Client) repository.AccountNotification {
	return &AccountNotification{
		tableName: "account_notification",
		database:  database,
	}
}

func (r *AccountNotification) Create(notification *entity.AccountNotification) error {
	_, err := r.database.Exec(context.Background(), `
		INSERT INTO "`+r.tableName+`" (id, ts, created_at, deleted_at, kind, level, status, message, metadata, artifact_id, cluster_id, account_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`, notification.ID, notification.TS, notification.CreatedAt, notification.DeletedAt, notification.Kind, notification.Level, notification.Status, notification.Message, notification.Metadata, notification.ArtifactID, notification.ClusterID, notification.AccountID)
	return err
}

func (r *AccountNotification) Save(notification *entity.AccountNotification) error {
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`"
		SET ts = $2, deleted_at = $3, kind = $4, level = $5, status = $6, message = $7, metadata = $8, artifact_id = $9
		WHERE id = $1
	`, notification.ID, notification.TS, notification.DeletedAt, notification.Kind, notification.Level, notification.Status, notification.Message, notification.Metadata, notification.ArtifactID)
	return err
}

func (r *AccountNotification) Delete(id uuid.UUID) error {
	now := time.Now()
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`" SET deleted_at = $2, ts = $2 WHERE id = $1
	`, id, now)
	return err
}

func (r *AccountNotification) FindById(id uuid.UUID) (*entity.AccountNotification, error) {
	row := r.database.QueryRow(context.Background(), `
		SELECT id, ts, created_at, deleted_at, kind, level, status, message, metadata, artifact_id, cluster_id, account_id
		FROM "`+r.tableName+`"
		WHERE id = $1
	`, id)

	var notification entity.AccountNotification
	err := row.Scan(
		&notification.ID,
		&notification.TS,
		&notification.CreatedAt,
		&notification.DeletedAt,
		&notification.Kind,
		&notification.Level,
		&notification.Status,
		&notification.Message,
		&notification.Metadata,
		&notification.ArtifactID,
		&notification.ClusterID,
		&notification.AccountID,
	)
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *AccountNotification) FindByAccountId(accountId uuid.UUID) ([]*entity.AccountNotification, error) {
	rows, err := r.database.Query(context.Background(), `
		SELECT id, ts, created_at, deleted_at, kind, level, status, message, metadata, artifact_id, cluster_id, account_id
		FROM "`+r.tableName+`"
		WHERE account_id = $1
	`, accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*entity.AccountNotification
	for rows.Next() {
		var notification entity.AccountNotification
		err := rows.Scan(
			&notification.ID,
			&notification.TS,
			&notification.CreatedAt,
			&notification.DeletedAt,
			&notification.Kind,
			&notification.Level,
			&notification.Status,
			&notification.Message,
			&notification.Metadata,
			&notification.ArtifactID,
			&notification.ClusterID,
			&notification.AccountID,
		)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}
	return notifications, nil
}

func (r *AccountNotification) FindByClusterId(clusterId uuid.UUID) ([]*entity.AccountNotification, error) {
	rows, err := r.database.Query(context.Background(), `
		SELECT id, ts, created_at, deleted_at, kind, level, status, message, metadata, artifact_id, cluster_id, account_id
		FROM "`+r.tableName+`"
		WHERE cluster_id = $1
	`, clusterId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*entity.AccountNotification
	for rows.Next() {
		var notification entity.AccountNotification
		err := rows.Scan(
			&notification.ID,
			&notification.TS,
			&notification.CreatedAt,
			&notification.DeletedAt,
			&notification.Kind,
			&notification.Level,
			&notification.Status,
			&notification.Message,
			&notification.Metadata,
			&notification.ArtifactID,
			&notification.ClusterID,
			&notification.AccountID,
		)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}
	return notifications, nil
}

func (r *AccountNotification) CountUnreadByAccountId(accountId uuid.UUID) (int64, error) {
	row := r.database.QueryRow(context.Background(), `
		SELECT count(*)
		FROM "`+r.tableName+`"
		WHERE account_id = $1 AND status = $2
	`, accountId, enum.AccountNotificationStatus_ACTIVE)

	var count int64
	err := row.Scan(&count)
	return count, err
}

func (r *AccountNotification) MarkAsReadByAccountId(accountId uuid.UUID) error {
	now := time.Now()
	_, err := r.database.Exec(context.Background(), `
		UPDATE "`+r.tableName+`"
		SET status = $2, ts = $3
		WHERE account_id = $1 AND status = $4	
	`, accountId, enum.AccountNotificationStatus_DISMISSED, now, enum.AccountNotificationStatus_ACTIVE)
	return err
}
