// domain/tenant/event/cluster.go
package event

import (
	"src/domain"

	"github.com/google/uuid"
)

type Cluster string

const (
	Cluster_CREATED         Cluster = "cluster.created"
	Cluster_PROMOTED        Cluster = "cluster.promoted"
	Cluster_MIGRATED        Cluster = "cluster.migrated"
	Cluster_DELETED         Cluster = "cluster.deleted"
	ClusterAgent_REGISTERED Cluster = "cluster.agent_registered"
	ClusterAgent_HEARTBEAT  Cluster = "cluster.agent_heartbeat"
	ClusterRuntime_DEPLOYED Cluster = "cluster.runtime_deployed"
	ClusterRuntime_UPDATED  Cluster = "cluster.runtime_updated"
	ClusterRuntime_DELETED  Cluster = "cluster.runtime_deleted"
)

type ClusterCreatedPayload struct {
	ClusterID      uuid.UUID `json:"cluster_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	Name           string    `json:"name"`
	Namespace      string    `json:"namespace"`
}

func ClusterCreated(
	clusterID uuid.UUID,
	organizationID uuid.UUID,
	name string,
	namespace string,
) domain.Event[Cluster] {
	return domain.NewEvent(
		Cluster_CREATED,
		ClusterCreatedPayload{
			ClusterID:      clusterID,
			OrganizationID: organizationID,
			Name:           name,
			Namespace:      namespace,
		},
	)
}

type ClusterPromotedPayload struct {
	ClusterID uuid.UUID `json:"cluster_id"`
}

func ClusterPromoted(
	clusterID uuid.UUID,
) domain.Event[Cluster] {
	return domain.NewEvent(
		Cluster_PROMOTED,
		ClusterPromotedPayload{
			ClusterID: clusterID,
		},
	)
}

type ClusterMigratedPayload struct {
	ClusterID uuid.UUID `json:"cluster_id"`
}

func ClusterMigrated(
	clusterID uuid.UUID,
) domain.Event[Cluster] {
	return domain.NewEvent(
		Cluster_MIGRATED,
		ClusterMigratedPayload{
			ClusterID: clusterID,
		},
	)
}

type ClusterDeletedPayload struct {
	ClusterID uuid.UUID `json:"cluster_id"`
}

func ClusterDeleted(
	clusterID uuid.UUID,
) domain.Event[Cluster] {
	return domain.NewEvent(
		Cluster_DELETED,
		ClusterDeletedPayload{
			ClusterID: clusterID,
		},
	)
}

// Nested entity: Agent
type ClusterAgentRegisteredPayload struct {
	AgentID   uuid.UUID `json:"agent_id"`
	ClusterID uuid.UUID `json:"cluster_id"`
	Version   string    `json:"version"`
}

func ClusterAgentRegistered(
	agentID uuid.UUID,
	clusterID uuid.UUID,
	version string,
) domain.Event[Cluster] {
	return domain.NewEvent(
		ClusterAgent_REGISTERED,
		ClusterAgentRegisteredPayload{
			AgentID:   agentID,
			ClusterID: clusterID,
			Version:   version,
		},
	)
}

type ClusterAgentHeartbeatPayload struct {
	AgentID   uuid.UUID `json:"agent_id"`
	ClusterID uuid.UUID `json:"cluster_id"`
}

func ClusterAgentHeartbeat(
	agentID uuid.UUID,
	clusterID uuid.UUID,
) domain.Event[Cluster] {
	return domain.NewEvent(
		ClusterAgent_HEARTBEAT,
		ClusterAgentHeartbeatPayload{
			AgentID:   agentID,
			ClusterID: clusterID,
		},
	)
}

// Nested entity: Runtime
type ClusterRuntimeDeployedPayload struct {
	RuntimeID uuid.UUID `json:"runtime_id"`
	ClusterID uuid.UUID `json:"cluster_id"`
}

func ClusterRuntimeDeployed(
	runtimeID uuid.UUID,
	clusterID uuid.UUID,
) domain.Event[Cluster] {
	return domain.NewEvent(
		ClusterRuntime_DEPLOYED,
		ClusterRuntimeDeployedPayload{
			RuntimeID: runtimeID,
			ClusterID: clusterID,
		},
	)
}

type ClusterRuntimeUpdatedPayload struct {
	RuntimeID uuid.UUID `json:"runtime_id"`
	ClusterID uuid.UUID `json:"cluster_id"`
}

func ClusterRuntimeUpdated(
	runtimeID uuid.UUID,
	clusterID uuid.UUID,
) domain.Event[Cluster] {
	return domain.NewEvent(
		ClusterRuntime_UPDATED,
		ClusterRuntimeUpdatedPayload{
			RuntimeID: runtimeID,
			ClusterID: clusterID,
		},
	)
}

type ClusterRuntimeDeletedPayload struct {
	RuntimeID uuid.UUID `json:"runtime_id"`
	ClusterID uuid.UUID `json:"cluster_id"`
}

func ClusterRuntimeDeleted(
	runtimeID uuid.UUID,
	clusterID uuid.UUID,
) domain.Event[Cluster] {
	return domain.NewEvent(
		ClusterRuntime_DELETED,
		ClusterRuntimeDeletedPayload{
			RuntimeID: runtimeID,
			ClusterID: clusterID,
		},
	)
}

var ClusterMapper = domain.NewEventMapper[Cluster]().
	Decoder(Cluster_CREATED, domain.JSONDecoder[ClusterCreatedPayload]()).
	Decoder(Cluster_PROMOTED, domain.JSONDecoder[ClusterPromotedPayload]()).
	Decoder(Cluster_MIGRATED, domain.JSONDecoder[ClusterMigratedPayload]()).
	Decoder(Cluster_DELETED, domain.JSONDecoder[ClusterDeletedPayload]()).
	Decoder(ClusterAgent_REGISTERED, domain.JSONDecoder[ClusterAgentRegisteredPayload]()).
	Decoder(ClusterAgent_HEARTBEAT, domain.JSONDecoder[ClusterAgentHeartbeatPayload]()).
	Decoder(ClusterRuntime_DEPLOYED, domain.JSONDecoder[ClusterRuntimeDeployedPayload]()).
	Decoder(ClusterRuntime_UPDATED, domain.JSONDecoder[ClusterRuntimeUpdatedPayload]()).
	Decoder(ClusterRuntime_DELETED, domain.JSONDecoder[ClusterRuntimeDeletedPayload]())
