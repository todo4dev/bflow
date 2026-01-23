// domain/event/pipeline.go
package event

import (
	"src/domain"

	"github.com/google/uuid"
)

const (
	Pipeline_REQUESTED       = "pipeline.requested"
	Pipeline_STARTED         = "pipeline.started"
	Pipeline_COMPLETED       = "pipeline.completed"
	Pipeline_FAILED          = "pipeline.failed"
	Pipeline_CANCELED        = "pipeline.canceled"
	Pipeline_RETRIED         = "pipeline.retried"
	PipelineAction_STARTED   = "pipeline.action_started"
	PipelineAction_COMPLETED = "pipeline.action_completed"
	PipelineAction_FAILED    = "pipeline.action_failed"
	PipelineStage_STARTED    = "pipeline.stage_started"
	PipelineStage_COMPLETED  = "pipeline.stage_completed"
	PipelineStage_FAILED     = "pipeline.stage_failed"
)

type PipelineRequestedPayload struct {
	PipelineID         uuid.UUID  `json:"pipeline_id"`
	RuntimeID          uuid.UUID  `json:"runtime_id"`
	RequesterAccountID uuid.UUID  `json:"requester_account_id"`
	Kind               string     `json:"kind"`
	TargetReleaseID    *uuid.UUID `json:"target_release_id"`
}

func PipelineRequested(pipelineID, runtimeID, requesterAccountID uuid.UUID, kind string, targetReleaseID *uuid.UUID) domain.Event {
	return domain.NewEvent(Pipeline_REQUESTED, PipelineRequestedPayload{
		PipelineID:         pipelineID,
		RuntimeID:          runtimeID,
		RequesterAccountID: requesterAccountID,
		Kind:               kind,
		TargetReleaseID:    targetReleaseID,
	})
}

type PipelineStartedPayload struct {
	PipelineID uuid.UUID `json:"pipeline_id"`
}

func PipelineStarted(pipelineID uuid.UUID) domain.Event {
	return domain.NewEvent(Pipeline_STARTED, PipelineStartedPayload{
		PipelineID: pipelineID,
	})
}

type PipelineCompletedPayload struct {
	PipelineID uuid.UUID `json:"pipeline_id"`
}

func PipelineCompleted(pipelineID uuid.UUID) domain.Event {
	return domain.NewEvent(Pipeline_COMPLETED, PipelineCompletedPayload{
		PipelineID: pipelineID,
	})
}

type PipelineFailedPayload struct {
	PipelineID   uuid.UUID `json:"pipeline_id"`
	ErrorMessage string    `json:"error_message"`
}

func PipelineFailed(pipelineID uuid.UUID, errorMessage string) domain.Event {
	return domain.NewEvent(Pipeline_FAILED, PipelineFailedPayload{
		PipelineID:   pipelineID,
		ErrorMessage: errorMessage,
	})
}

type PipelineCanceledPayload struct {
	PipelineID uuid.UUID `json:"pipeline_id"`
}

func PipelineCanceled(pipelineID uuid.UUID) domain.Event {
	return domain.NewEvent(Pipeline_CANCELED, PipelineCanceledPayload{
		PipelineID: pipelineID,
	})
}

type PipelineRetriedPayload struct {
	PipelineID         uuid.UUID `json:"pipeline_id"`
	PreviousPipelineID uuid.UUID `json:"previous_pipeline_id"`
}

func PipelineRetried(pipelineID, previousPipelineID uuid.UUID) domain.Event {
	return domain.NewEvent(Pipeline_RETRIED, PipelineRetriedPayload{
		PipelineID:         pipelineID,
		PreviousPipelineID: previousPipelineID,
	})
}

// Nested entity: Action
type PipelineActionStartedPayload struct {
	ActionID         uuid.UUID `json:"action_id"`
	PipelineID       uuid.UUID `json:"pipeline_id"`
	Kind             string    `json:"kind"`
	ExecutionAgentID uuid.UUID `json:"execution_agent_id"`
}

func PipelineActionStarted(actionID, pipelineID uuid.UUID, kind string, executionAgentID uuid.UUID) domain.Event {
	return domain.NewEvent(PipelineAction_STARTED, PipelineActionStartedPayload{
		ActionID:         actionID,
		PipelineID:       pipelineID,
		Kind:             kind,
		ExecutionAgentID: executionAgentID,
	})
}

type PipelineActionCompletedPayload struct {
	ActionID   uuid.UUID `json:"action_id"`
	PipelineID uuid.UUID `json:"pipeline_id"`
}

func PipelineActionCompleted(actionID, pipelineID uuid.UUID) domain.Event {
	return domain.NewEvent(PipelineAction_COMPLETED, PipelineActionCompletedPayload{
		ActionID:   actionID,
		PipelineID: pipelineID,
	})
}

type PipelineActionFailedPayload struct {
	ActionID     uuid.UUID `json:"action_id"`
	PipelineID   uuid.UUID `json:"pipeline_id"`
	ErrorMessage string    `json:"error_message"`
}

func PipelineActionFailed(actionID, pipelineID uuid.UUID, errorMessage string) domain.Event {
	return domain.NewEvent(PipelineAction_FAILED, PipelineActionFailedPayload{
		ActionID:     actionID,
		PipelineID:   pipelineID,
		ErrorMessage: errorMessage,
	})
}

type PipelineStageStartedPayload struct {
	StageID  uuid.UUID `json:"stage_id"`
	ActionID uuid.UUID `json:"action_id"`
	Name     string    `json:"name"`
	Position int       `json:"position"`
}

func PipelineStageStarted(stageID, actionID uuid.UUID, name string, position int) domain.Event {
	return domain.NewEvent(PipelineStage_STARTED, PipelineStageStartedPayload{
		StageID:  stageID,
		ActionID: actionID,
		Name:     name,
		Position: position,
	})
}

type PipelineStageCompletedPayload struct {
	StageID  uuid.UUID `json:"stage_id"`
	ActionID uuid.UUID `json:"action_id"`
	Summary  string    `json:"summary"`
}

func PipelineStageCompleted(stageID, actionID uuid.UUID, summary string) domain.Event {
	return domain.NewEvent(PipelineStage_COMPLETED, PipelineStageCompletedPayload{
		StageID:  stageID,
		ActionID: actionID,
		Summary:  summary,
	})
}

type PipelineStageFailedPayload struct {
	StageID  uuid.UUID `json:"stage_id"`
	ActionID uuid.UUID `json:"action_id"`
	Summary  string    `json:"summary"`
}

func PipelineStageFailed(stageID, actionID uuid.UUID, summary string) domain.Event {
	return domain.NewEvent(PipelineStage_FAILED, PipelineStageFailedPayload{
		StageID:  stageID,
		ActionID: actionID,
		Summary:  summary,
	})
}

var PipelineMapper = domain.NewEventMapper().
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
