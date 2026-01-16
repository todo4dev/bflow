// domain/deployment/enum/agent.go
package enum

type AgentStatus string

const (
	AgentStatus_PENDING  AgentStatus = "PENDING"
	AgentStatus_ONLINE   AgentStatus = "ONLINE"
	AgentStatus_OFFLINE  AgentStatus = "OFFLINE"
	AgentStatus_DISABLED AgentStatus = "DISABLED"
	AgentStatus_DRAINING AgentStatus = "DRAINING"
)

type AgentAuthType string

const (
	AgentAuthType_SECRET AgentAuthType = "SECRET"
	AgentAuthType_KEY    AgentAuthType = "KEY"
)
