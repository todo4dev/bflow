// domain/uow.go
package domain

import "src/domain/repository"

type Repository interface {
	Account() repository.Account
	AccountActivity() repository.AccountActivity
	AccountCertificate() repository.AccountCertificate
	AccountCredential() repository.AccountCredential
	AccountNotification() repository.AccountNotification
	AccountPreference() repository.AccountPreference
	AccountProfile() repository.AccountProfile
	Artifact() repository.Artifact
	ArtifactRelease() repository.ArtifactRelease
	Cluster() repository.Cluster
	ClusterAgent() repository.ClusterAgent
	ClusterAgentEnrollment() repository.ClusterAgentEnrollment
	ClusterRuntime() repository.ClusterRuntime
	Document() repository.Document
	DocumentSignature() repository.DocumentSignature
	Organization() repository.Organization
	OrganizationInvite() repository.OrganizationInvite
	OrganizationMembership() repository.OrganizationMembership
	Pipeline() repository.Pipeline
	PipelineAction() repository.PipelineAction
	PipelineActionStage() repository.PipelineActionStage
	Plan() repository.Plan
	Subscription() repository.Subscription
	SubscriptionInvoice() repository.SubscriptionInvoice
	SubscriptionPayment() repository.SubscriptionPayment
}

type Uow interface {
	Repository

	Do(func(t Repository) error) error
}
