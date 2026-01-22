// domain/repository/document_signature.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type DocumentSignature interface {
	Create(signature *entity.DocumentSignature) error
	Save(signature *entity.DocumentSignature) error
	FindById(id uuid.UUID) (*entity.DocumentSignature, error)
	FindByDocumentId(documentId uuid.UUID) ([]*entity.DocumentSignature, error)
}
