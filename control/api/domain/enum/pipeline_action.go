// domain/enum/pipeline_action.go
package enum

type PipelineActionKind string

const (
	PipelineActionKind_DEPLOY   PipelineActionKind = "DEPLOY"
	PipelineActionKind_CONFIG   PipelineActionKind = "CONFIG"
	PipelineActionKind_ROLLBACK PipelineActionKind = "ROLLBACK"
	PipelineActionKind_DESTROY  PipelineActionKind = "DESTROY"
)

type PipelineActionStatus string

const (
	PipelineActionStatus_PENDING  PipelineActionStatus = "PENDING"
	PipelineActionStatus_RUNNING  PipelineActionStatus = "RUNNING"
	PipelineActionStatus_SUCCESS  PipelineActionStatus = "SUCCESS"
	PipelineActionStatus_FAILURE  PipelineActionStatus = "FAILURE"
	PipelineActionStatus_CANCELED PipelineActionStatus = "CANCELED"
)

