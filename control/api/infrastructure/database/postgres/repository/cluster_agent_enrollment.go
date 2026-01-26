// infrastructure/database/pgx/repository/cluster_agent_enrollment.go
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

type ClusterAgentEnrollment struct {
	client database.Client
}

func NewClusterAgentEnrollment(client database.Client) *ClusterAgentEnrollment {
	return &ClusterAgentEnrollment{client: client}
}

var _ repository.ClusterAgentEnrollment = (*ClusterAgentEnrollment)(nil)

func (r *ClusterAgentEnrollment) Create(t *entity.ClusterAgentEnrollment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO "cluster_agent_enrollment" (
			"id", "ts", "created_at", "token", "expires_at", "used", "cluster_id", "creator_account_id"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.CreatedAt, t.Token, t.ExpiresAt, t.Used, t.ClusterID, t.CreatorAccountID,
	)
	return err
}

func (r *ClusterAgentEnrollment) Save(t *entity.ClusterAgentEnrollment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE "cluster_agent_enrollment" SET
			"ts" = $2, "token" = $3, "expires_at" = $4, "used" = $5
		WHERE "id" = $1
	`
	_, err := r.client.Exec(ctx, query,
		t.ID, t.TS, t.Token, t.ExpiresAt, t.Used,
	)
	return err
}

func (r *ClusterAgentEnrollment) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM "cluster_agent_enrollment" WHERE "id" = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}

func (r *ClusterAgentEnrollment) FindById(id uuid.UUID) (*entity.ClusterAgentEnrollment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "token", "expires_at", "used", "cluster_id", "creator_account_id"
		FROM "cluster_agent_enrollment"
		WHERE "id" = $1
	`
	var t entity.ClusterAgentEnrollment
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.Token, &t.ExpiresAt, &t.Used, &t.ClusterID, &t.CreatorAccountID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *ClusterAgentEnrollment) FindByToken(token string) (*entity.ClusterAgentEnrollment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT
			"id", "ts", "created_at", "token", "expires_at", "used", "cluster_id", "creator_account_id"
		FROM "cluster_agent_enrollment"
		WHERE "token" = $1
	`
	var t entity.ClusterAgentEnrollment
	if err := r.client.QueryRow(ctx, query, token).Scan(
		&t.ID, &t.TS, &t.CreatedAt, &t.Token, &t.ExpiresAt, &t.Used, &t.ClusterID, &t.CreatorAccountID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}
