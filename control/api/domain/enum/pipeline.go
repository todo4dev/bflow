// domain/enum/pipeline.go
package enum

type PipelineKind string

const (
	PipelineKind_BUILD  PipelineKind = "BUILD"
	PipelineKind_DEPLOY PipelineKind = "DEPLOY"
)

type PipelineStatus string

const (
	PipelineStatus_PENDING   PipelineStatus = "PENDING"
	PipelineStatus_SUCCEEDED PipelineStatus = "SUCCEEDED"
	PipelineStatus_FAILED    PipelineStatus = "FAILED"
	PipelineStatus_CANCELED  PipelineStatus = "CANCELED"
)

