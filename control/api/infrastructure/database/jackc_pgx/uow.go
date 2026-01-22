// infrastructure/database/jackc_pgx/uow.go
package jackc_pgx

import (
	"context"
	"src/domain"
	"src/domain/repository"
	impl "src/infrastructure/database/jackc_pgx/repository"
	"src/port/database"
)

type Uow struct {
	client                 database.Client
	accountActivity        repository.AccountActivity
	accountCertificate     repository.AccountCertificate
	accountCredential      repository.AccountCredential
	accountNotification    repository.AccountNotification
	accountPreference      repository.AccountPreference
	accountProfile         repository.AccountProfile
	account                repository.Account
	artifactRelease        repository.ArtifactRelease
	artifact               repository.Artifact
	clusterAgentEnrollment repository.ClusterAgentEnrollment
	clusterAgent           repository.ClusterAgent
	clusterRuntime         repository.ClusterRuntime
	cluster                repository.Cluster
	documentSignature      repository.DocumentSignature
	document               repository.Document
	organizationInvite     repository.OrganizationInvite
	organizationMembership repository.OrganizationMembership
	organization           repository.Organization
	pipelineActionStage    repository.PipelineActionStage
	pipelineAction         repository.PipelineAction
	pipeline               repository.Pipeline
	plan                   repository.Plan
	subscriptionInvoice    repository.SubscriptionInvoice
	subscriptionPayment    repository.SubscriptionPayment
	subscription           repository.Subscription
}

var _ domain.Uow = (*Uow)(nil)

func NewUow(client database.Client) *Uow {
	return &Uow{
		client:                 client,
		accountActivity:        impl.NewAccountActivity(client),
		accountCertificate:     impl.NewAccountCertificate(client),
		accountCredential:      impl.NewAccountCredential(client),
		accountNotification:    impl.NewAccountNotification(client),
		accountPreference:      impl.NewAccountPreference(client),
		accountProfile:         impl.NewAccountProfile(client),
		account:                impl.NewAccount(client),
		artifactRelease:        impl.NewArtifactRelease(client),
		artifact:               impl.NewArtifact(client),
		clusterAgentEnrollment: impl.NewClusterAgentEnrollment(client),
		clusterAgent:           impl.NewClusterAgent(client),
		clusterRuntime:         impl.NewClusterRuntime(client),
		cluster:                impl.NewCluster(client),
		documentSignature:      impl.NewDocumentSignature(client),
		document:               impl.NewDocument(client),
		organizationInvite:     impl.NewOrganizationInvite(client),
		organizationMembership: impl.NewOrganizationMembership(client),
		organization:           impl.NewOrganization(client),
		pipelineActionStage:    impl.NewPipelineActionStage(client),
		pipelineAction:         impl.NewPipelineAction(client),
		pipeline:               impl.NewPipeline(client),
		plan:                   impl.NewPlan(client),
		subscriptionInvoice:    impl.NewSubscriptionInvoice(client),
		subscriptionPayment:    impl.NewSubscriptionPayment(client),
		subscription:           impl.NewSubscription(client),
	}
}

func (u *Uow) AccountActivity() repository.AccountActivity {
	return u.accountActivity
}

func (u *Uow) AccountCertificate() repository.AccountCertificate {
	return u.accountCertificate
}

func (u *Uow) AccountCredential() repository.AccountCredential {
	return u.accountCredential
}

func (u *Uow) AccountNotification() repository.AccountNotification {
	return u.accountNotification
}

func (u *Uow) AccountPreference() repository.AccountPreference {
	return u.accountPreference
}

func (u *Uow) AccountProfile() repository.AccountProfile {
	return u.accountProfile
}

func (u *Uow) Account() repository.Account {
	return u.account
}

func (u *Uow) ArtifactRelease() repository.ArtifactRelease {
	return u.artifactRelease
}

func (u *Uow) Artifact() repository.Artifact {
	return u.artifact
}

func (u *Uow) ClusterAgentEnrollment() repository.ClusterAgentEnrollment {
	return u.clusterAgentEnrollment
}

func (u *Uow) ClusterAgent() repository.ClusterAgent {
	return u.clusterAgent
}

func (u *Uow) ClusterRuntime() repository.ClusterRuntime {
	return u.clusterRuntime
}

func (u *Uow) Cluster() repository.Cluster {
	return u.cluster
}

func (u *Uow) DocumentSignature() repository.DocumentSignature {
	return u.documentSignature
}

func (u *Uow) Document() repository.Document {
	return u.document
}

func (u *Uow) OrganizationInvite() repository.OrganizationInvite {
	return u.organizationInvite
}

func (u *Uow) OrganizationMembership() repository.OrganizationMembership {
	return u.organizationMembership
}

func (u *Uow) Organization() repository.Organization {
	return u.organization
}

func (u *Uow) PipelineActionStage() repository.PipelineActionStage {
	return u.pipelineActionStage
}

func (u *Uow) PipelineAction() repository.PipelineAction {
	return u.pipelineAction
}

func (u *Uow) Pipeline() repository.Pipeline {
	return u.pipeline
}

func (u *Uow) Plan() repository.Plan {
	return u.plan
}

func (u *Uow) SubscriptionInvoice() repository.SubscriptionInvoice {
	return u.subscriptionInvoice
}

func (u *Uow) SubscriptionPayment() repository.SubscriptionPayment {
	return u.subscriptionPayment
}

func (u *Uow) Subscription() repository.Subscription {
	return u.subscription
}

func (u *Uow) Do(fn func(t domain.Repository) error) error {
	return u.client.Transaction(context.Background(), func(tx database.Client) error {
		return fn(NewUow(tx))
	})
}
