// domain/event/pipeline.go
package event

import (
	"src/domain"

	"github.com/google/uuid"
)

type Pipeline string

const (
	Pipeline_REQUESTED       Pipeline = "pipeline.requested"
	Pipeline_STARTED         Pipeline = "pipeline.started"
	Pipeline_COMPLETED       Pipeline = "pipeline.completed"
	Pipeline_FAILED          Pipeline = "pipeline.failed"
	Pipeline_CANCELED        Pipeline = "pipeline.canceled"
	Pipeline_RETRIED         Pipeline = "pipeline.retried"
	PipelineAction_STARTED   Pipeline = "pipeline.action_started"
	PipelineAction_COMPLETED Pipeline = "pipeline.action_completed"
	PipelineAction_FAILED    Pipeline = "pipeline.action_failed"
	PipelineStage_STARTED    Pipeline = "pipeline.stage_started"
	PipelineStage_COMPLETED  Pipeline = "pipeline.stage_completed"
	PipelineStage_FAILED     Pipeline = "pipeline.stage_failed"
)

type PipelineRequestedPayload struct {
	PipelineID         uuid.UUID  `json:"pipeline_id"`
	Kind               string     `json:"kind"`
	RuntimeID          uuid.UUID  `json:"runtime_id"`
	TargetReleaseID    *uuid.UUID `json:"target_release_id"`
	RequesterAccountID uuid.UUID  `json:"requester_account_id"`
}

func PipelineRequested(
	pipelineID uuid.UUID,
	kind string,
	runtimeID uuid.UUID,
	targetReleaseID *uuid.UUID,
	requesterAccountID uuid.UUID,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		Pipeline_REQUESTED,
		PipelineRequestedPayload{
			PipelineID:         pipelineID,
			Kind:               kind,
			RuntimeID:          runtimeID,
			TargetReleaseID:    targetReleaseID,
			RequesterAccountID: requesterAccountID,
		},
	)
}

type PipelineStartedPayload struct {
	PipelineID uuid.UUID `json:"pipeline_id"`
}

func PipelineStarted(
	pipelineID uuid.UUID,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		Pipeline_STARTED,
		PipelineStartedPayload{
			PipelineID: pipelineID,
		},
	)
}

type PipelineCompletedPayload struct {
	PipelineID uuid.UUID `json:"pipeline_id"`
}

func PipelineCompleted(
	pipelineID uuid.UUID,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		Pipeline_COMPLETED,
		PipelineCompletedPayload{
			PipelineID: pipelineID,
		},
	)
}

type PipelineFailedPayload struct {
	PipelineID   uuid.UUID `json:"pipeline_id"`
	ErrorMessage string    `json:"error_message"`
}

func PipelineFailed(
	pipelineID uuid.UUID,
	errorMessage string,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		Pipeline_FAILED,
		PipelineFailedPayload{
			PipelineID:   pipelineID,
			ErrorMessage: errorMessage,
		},
	)
}

type PipelineCanceledPayload struct {
	PipelineID uuid.UUID `json:"pipeline_id"`
}

func PipelineCanceled(
	pipelineID uuid.UUID,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		Pipeline_CANCELED,
		PipelineCanceledPayload{
			PipelineID: pipelineID,
		},
	)
}

type PipelineRetriedPayload struct {
	PipelineID         uuid.UUID `json:"pipeline_id"`
	PreviousPipelineID uuid.UUID `json:"previous_pipeline_id"`
}

func PipelineRetried(
	pipelineID uuid.UUID,
	previousPipelineID uuid.UUID,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		Pipeline_RETRIED,
		PipelineRetriedPayload{
			PipelineID:         pipelineID,
			PreviousPipelineID: previousPipelineID,
		},
	)
}

// Nested entity: Action
type PipelineActionStartedPayload struct {
	ActionID         uuid.UUID `json:"action_id"`
	PipelineID       uuid.UUID `json:"pipeline_id"`
	Kind             string    `json:"kind"`
	ExecutionAgentID uuid.UUID `json:"execution_agent_id"`
}

func PipelineActionStarted(
	actionID uuid.UUID,
	pipelineID uuid.UUID,
	kind string,
	executionAgentID uuid.UUID,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		PipelineAction_STARTED,
		PipelineActionStartedPayload{
			ActionID:         actionID,
			PipelineID:       pipelineID,
			Kind:             kind,
			ExecutionAgentID: executionAgentID,
		},
	)
}

type PipelineActionCompletedPayload struct {
	ActionID   uuid.UUID `json:"action_id"`
	PipelineID uuid.UUID `json:"pipeline_id"`
}

func PipelineActionCompleted(
	actionID uuid.UUID,
	pipelineID uuid.UUID,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		PipelineAction_COMPLETED,
		PipelineActionCompletedPayload{
			ActionID:   actionID,
			PipelineID: pipelineID,
		},
	)
}

type PipelineActionFailedPayload struct {
	ActionID     uuid.UUID `json:"action_id"`
	PipelineID   uuid.UUID `json:"pipeline_id"`
	ErrorMessage string    `json:"error_message"`
}

func PipelineActionFailed(
	actionID uuid.UUID,
	pipelineID uuid.UUID,
	errorMessage string,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		PipelineAction_FAILED,
		PipelineActionFailedPayload{
			ActionID:     actionID,
			PipelineID:   pipelineID,
			ErrorMessage: errorMessage,
		},
	)
}

// Nested entity: Stage
type PipelineStageStartedPayload struct {
	StageID  uuid.UUID `json:"stage_id"`
	ActionID uuid.UUID `json:"action_id"`
	Name     string    `json:"name"`
	Position int       `json:"position"`
}

func PipelineStageStarted(
	stageID uuid.UUID,
	actionID uuid.UUID,
	name string,
	position int,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		PipelineStage_STARTED,
		PipelineStageStartedPayload{
			StageID:  stageID,
			ActionID: actionID,
			Name:     name,
			Position: position,
		},
	)
}

type PipelineStageCompletedPayload struct {
	StageID  uuid.UUID `json:"stage_id"`
	ActionID uuid.UUID `json:"action_id"`
	Summary  string    `json:"summary"`
}

func PipelineStageCompleted(
	stageID uuid.UUID,
	actionID uuid.UUID,
	summary string,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		PipelineStage_COMPLETED,
		PipelineStageCompletedPayload{
			StageID:  stageID,
			ActionID: actionID,
			Summary:  summary,
		},
	)
}

type PipelineStageFailedPayload struct {
	StageID  uuid.UUID `json:"stage_id"`
	ActionID uuid.UUID `json:"action_id"`
	Summary  string    `json:"summary"`
}

func PipelineStageFailed(
	stageID uuid.UUID,
	actionID uuid.UUID,
	summary string,
) domain.Event[Pipeline] {
	return domain.NewEvent(
		PipelineStage_FAILED,
		PipelineStageFailedPayload{
			StageID:  stageID,
			ActionID: actionID,
			Summary:  summary,
		},
	)
}

var PipelineMapper = domain.NewEventMapper[Pipeline]().
	Decoder(Pipeline_REQUESTED, domain.JSONDecoder[PipelineRequestedPayload]()).
	Decoder(Pipeline_STARTED, domain.JSONDecoder[PipelineStartedPayload]()).
	Decoder(Pipeline_COMPLETED, domain.JSONDecoder[PipelineCompletedPayload]()).
	Decoder(Pipeline_FAILED, domain.JSONDecoder[PipelineFailedPayload]()).
	Decoder(Pipeline_CANCELED, domain.JSONDecoder[PipelineCanceledPayload]()).
	Decoder(Pipeline_RETRIED, domain.JSONDecoder[PipelineRetriedPayload]()).
	Decoder(PipelineAction_STARTED, domain.JSONDecoder[PipelineActionStartedPayload]()).
	Decoder(PipelineAction_COMPLETED, domain.JSONDecoder[PipelineActionCompletedPayload]()).
	Decoder(PipelineAction_FAILED, domain.JSONDecoder[PipelineActionFailedPayload]()).
	Decoder(PipelineStage_STARTED, domain.JSONDecoder[PipelineStageStartedPayload]()).
	Decoder(PipelineStage_COMPLETED, domain.JSONDecoder[PipelineStageCompletedPayload]()).
	Decoder(PipelineStage_FAILED, domain.JSONDecoder[PipelineStageFailedPayload]())

