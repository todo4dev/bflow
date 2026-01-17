// domain/deployment/enum/cluster_agent.go
package enum

type ClusterAgentStatus string

const (
	ClusterAgentStatus_PENDING  ClusterAgentStatus = "PENDING"
	ClusterAgentStatus_ONLINE   ClusterAgentStatus = "ONLINE"
	ClusterAgentStatus_OFFLINE  ClusterAgentStatus = "OFFLINE"
	ClusterAgentStatus_DISABLED ClusterAgentStatus = "DISABLED"
	ClusterAgentStatus_DRAINING ClusterAgentStatus = "DRAINING"
)

type ClusterAgentAuthType string

const (
	ClusterAgentAuthType_SECRET ClusterAgentAuthType = "SECRET"
	ClusterAgentAuthType_KEY    ClusterAgentAuthType = "KEY"
)
