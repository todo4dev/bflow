// domain/tenant/enum/cluster.go
package enum

type ClusterState string

const (
	ClusterState_ACTIVE    ClusterState = "ACTIVE"
	ClusterState_MIGRATING ClusterState = "MIGRATING"
	ClusterState_LEGACY    ClusterState = "LEGACY"
	ClusterState_DISABLED  ClusterState = "DISABLED"
)
