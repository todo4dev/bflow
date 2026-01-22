// domain/repository/artifact.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type Artifact interface {
	Create(artifact *entity.Artifact) error
	Save(artifact *entity.Artifact) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Artifact, error)
}
