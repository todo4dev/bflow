// domain/repository/artifact_release.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type ArtifactRelease interface {
	Create(release *entity.ArtifactRelease) error
	Save(release *entity.ArtifactRelease) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.ArtifactRelease, error)
	FindByArtifactId(artifactId uuid.UUID) ([]*entity.ArtifactRelease, error)
}
