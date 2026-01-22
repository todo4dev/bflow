// domain/repository/plan.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type Plan interface {
	Create(plan *entity.Plan) error
	Save(plan *entity.Plan) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Plan, error)
	FindByCode(code string) (*entity.Plan, error)
}
