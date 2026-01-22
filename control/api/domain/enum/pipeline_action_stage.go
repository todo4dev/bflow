// domain/enum/pipeline_action_stage.go
package enum

type PipelineActionStageStatus string

const (
	PipelineActionStageStatus_PENDING PipelineActionStageStatus = "PENDING"
	PipelineActionStageStatus_RUNNING PipelineActionStageStatus = "RUNNING"
	PipelineActionStageStatus_SUCCESS PipelineActionStageStatus = "SUCCESS"
	PipelineActionStageStatus_FAILED  PipelineActionStageStatus = "FAILED"
)

