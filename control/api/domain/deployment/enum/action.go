// domain/deployment/enum/action.go
// domain/identity/enum/action.go
package enum

type ActionKind string

const (
	ActionKind_DEPLOY   ActionKind = "DEPLOY"
	ActionKind_CONFIG   ActionKind = "CONFIG"
	ActionKind_ROLLBACK ActionKind = "ROLLBACK"
	ActionKind_DESTROY  ActionKind = "DESTROY"
)

type ActionStatus string

const (
	ActionStatus_PENDING  ActionStatus = "PENDING"
	ActionStatus_RUNNING  ActionStatus = "RUNNING"
	ActionStatus_SUCCESS  ActionStatus = "SUCCESS"
	ActionStatus_FAILURE  ActionStatus = "FAILURE"
	ActionStatus_CANCELED ActionStatus = "CANCELED"
)
