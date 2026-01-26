// infrastructure/database/postgres/postgres.go
package postgres

import (
	"fmt"
	"src/domain"
	"src/domain/repository"
	impl "src/infrastructure/database/postgres/repository"
	"src/port/database"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	config := Config{
		DSN: env.Get("API_DATABASE_POSTGRES_DSN", "postgres://user:pass@localhost:5432/bflow?sslmode=disable"),
	}
	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("database config validation failed: %w", err))
	}

	instance, err := NewClient(&config)
	if err != nil {
		panic(fmt.Errorf("failed to create database client: %w", err))
	}

	di.SingletonAs[database.Client](func() database.Client { return instance })
	di.SingletonAs[repository.AccountActivity](impl.NewAccountActivity)
	di.SingletonAs[repository.AccountCertificate](impl.NewAccountCertificate)
	di.SingletonAs[repository.AccountCredential](impl.NewAccountCredential)
	di.SingletonAs[repository.AccountNotification](impl.NewAccountNotification)
	di.SingletonAs[repository.AccountPreference](impl.NewAccountPreference)
	di.SingletonAs[repository.AccountProfile](impl.NewAccountProfile)
	di.SingletonAs[repository.Account](impl.NewAccount)
	di.SingletonAs[repository.ArtifactRelease](impl.NewArtifactRelease)
	di.SingletonAs[repository.Artifact](impl.NewArtifact)
	di.SingletonAs[repository.ClusterAgentEnrollment](impl.NewClusterAgentEnrollment)
	di.SingletonAs[repository.ClusterAgent](impl.NewClusterAgent)
	di.SingletonAs[repository.ClusterRuntime](impl.NewClusterRuntime)
	di.SingletonAs[repository.Cluster](impl.NewCluster)
	di.SingletonAs[repository.DocumentSignature](impl.NewDocumentSignature)
	di.SingletonAs[repository.Document](impl.NewDocument)
	di.SingletonAs[repository.OrganizationInvite](impl.NewOrganizationInvite)
	di.SingletonAs[repository.OrganizationMembership](impl.NewOrganizationMembership)
	di.SingletonAs[repository.Organization](impl.NewOrganization)
	di.SingletonAs[repository.PipelineActionStage](impl.NewPipelineActionStage)
	di.SingletonAs[repository.PipelineAction](impl.NewPipelineAction)
	di.SingletonAs[repository.Pipeline](impl.NewPipeline)
	di.SingletonAs[repository.Plan](impl.NewPlan)
	di.SingletonAs[repository.SubscriptionInvoice](impl.NewSubscriptionInvoice)
	di.SingletonAs[repository.SubscriptionPayment](impl.NewSubscriptionPayment)
	di.SingletonAs[repository.Subscription](impl.NewSubscription)
	di.SingletonAs[domain.Uow](NewUow)
}
