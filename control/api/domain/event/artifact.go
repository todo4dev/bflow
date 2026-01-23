// domain/event/artifact.go
package event

import (
	"src/domain"

	"github.com/google/uuid"
)

const (
	Artifact_UPLOADED         = "artifact.uploaded"
	ArtifactRelease_PUBLISHED = "artifact.release_published"
	ArtifactRelease_PROMOTED  = "artifact.release_promoted"
)

type ArtifactUploadedPayload struct {
	ArtifactID uuid.UUID `json:"artifact_id"`
	Kind       string    `json:"kind"`
	Name       string    `json:"name"`
}

func ArtifactUploaded(artifactID uuid.UUID, kind string, name string) domain.Event {
	return domain.NewEvent(Artifact_UPLOADED, ArtifactUploadedPayload{
		ArtifactID: artifactID,
		Kind:       kind,
		Name:       name,
	})
}

// Nested entity: Release
type ArtifactReleasePublishedPayload struct {
	ReleaseID  uuid.UUID `json:"release_id"`
	ArtifactID uuid.UUID `json:"artifact_id"`
	Version    string    `json:"version"`
	Channel    string    `json:"channel"`
}

func ArtifactReleasePublished(releaseID, artifactID uuid.UUID, version, channel string) domain.Event {
	return domain.NewEvent(ArtifactRelease_PUBLISHED, ArtifactReleasePublishedPayload{
		ReleaseID:  releaseID,
		ArtifactID: artifactID,
		Version:    version,
		Channel:    channel,
	})
}

type ArtifactReleasePromotedPayload struct {
	ReleaseID  uuid.UUID `json:"release_id"`
	ArtifactID uuid.UUID `json:"artifact_id"`
	Version    string    `json:"version"`
}

func ArtifactReleasePromoted(releaseID, artifactID uuid.UUID, version string) domain.Event {
	return domain.NewEvent(ArtifactRelease_PROMOTED, ArtifactReleasePromotedPayload{
		ReleaseID:  releaseID,
		ArtifactID: artifactID,
		Version:    version,
	})
}

var ArtifactMapper = domain.NewEventMapper().
	Decoder(Artifact_UPLOADED, domain.JSONDecoder[ArtifactUploadedPayload]()).
	Decoder(ArtifactRelease_PUBLISHED, domain.JSONDecoder[ArtifactReleasePublishedPayload]()).
	Decoder(ArtifactRelease_PROMOTED, domain.JSONDecoder[ArtifactReleasePromotedPayload]())
