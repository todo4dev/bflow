// domain/event/document.go
package event

import (
	"src/domain"

	"github.com/google/uuid"
)

const (
	Document_CREATED            = "document.created"
	Document_DELETED            = "document.deleted"
	Document_REPLACED           = "document.replaced"
	DocumentSignature_REQUESTED = "document.signature_requested"
	DocumentSignature_SIGNED    = "document.signature_signed"
	DocumentSignature_FAILED    = "document.signature_failed"
)

type DocumentCreatedPayload struct {
	DocumentID     uuid.UUID  `json:"document_id"`
	Kind           string     `json:"kind"`
	Title          string     `json:"title"`
	StorageKey     string     `json:"storage_key"`
	FileSize       int64      `json:"file_size"`
	ClusterID      *uuid.UUID `json:"cluster_id"`
	OrganizationID *uuid.UUID `json:"organization_id"`
}

func DocumentCreated(documentID uuid.UUID, kind, title, storageKey string, fileSize int64, clusterID, organizationID *uuid.UUID) domain.Event {
	return domain.NewEvent(Document_CREATED, DocumentCreatedPayload{
		DocumentID:     documentID,
		Kind:           kind,
		Title:          title,
		StorageKey:     storageKey,
		FileSize:       fileSize,
		ClusterID:      clusterID,
		OrganizationID: organizationID,
	})
}

type DocumentDeletedPayload struct {
	DocumentID uuid.UUID `json:"document_id"`
}

func DocumentDeleted(documentID uuid.UUID) domain.Event {
	return domain.NewEvent(Document_DELETED, DocumentDeletedPayload{
		DocumentID: documentID,
	})
}

type DocumentReplacedPayload struct {
	NewDocumentID uuid.UUID `json:"new_document_id"`
	OldDocumentID uuid.UUID `json:"old_document_id"`
}

func DocumentReplaced(newDocumentID uuid.UUID, oldDocumentID uuid.UUID) domain.Event {
	return domain.NewEvent(Document_REPLACED, DocumentReplacedPayload{
		NewDocumentID: newDocumentID,
		OldDocumentID: oldDocumentID,
	})
}

type DocumentSignatureRequestedPayload struct {
	SignatureID uuid.UUID `json:"signature_id"`
	DocumentID  uuid.UUID `json:"document_id"`
	AccountID   uuid.UUID `json:"account_id"`
}

func DocumentSignatureRequested(signatureID, documentID, accountID uuid.UUID) domain.Event {
	return domain.NewEvent(DocumentSignature_REQUESTED, DocumentSignatureRequestedPayload{
		SignatureID: signatureID,
		DocumentID:  documentID,
		AccountID:   accountID,
	})
}

type DocumentSignatureSignedPayload struct {
	SignatureID   uuid.UUID `json:"signature_id"`
	DocumentID    uuid.UUID `json:"document_id"`
	AccountID     uuid.UUID `json:"account_id"`
	CertificateID uuid.UUID `json:"certificate_id"`
}

func DocumentSignatureSigned(signatureID, documentID, accountID, certificateID uuid.UUID) domain.Event {
	return domain.NewEvent(DocumentSignature_SIGNED, DocumentSignatureSignedPayload{
		SignatureID:   signatureID,
		DocumentID:    documentID,
		AccountID:     accountID,
		CertificateID: certificateID,
	})
}

type DocumentSignatureFailedPayload struct {
	SignatureID   uuid.UUID `json:"signature_id"`
	DocumentID    uuid.UUID `json:"document_id"`
	AccountID     uuid.UUID `json:"account_id"`
	FailureReason string    `json:"failure_reason"`
}

func DocumentSignatureFailed(signatureID, documentID, accountID uuid.UUID, failureReason string) domain.Event {
	return domain.NewEvent(DocumentSignature_FAILED, DocumentSignatureFailedPayload{
		SignatureID:   signatureID,
		DocumentID:    documentID,
		AccountID:     accountID,
		FailureReason: failureReason,
	})
}

var DocumentMapper = domain.NewEventMapper().
	Decoder(Document_CREATED, domain.JSONDecoder[DocumentCreatedPayload]()).
	Decoder(Document_DELETED, domain.JSONDecoder[DocumentDeletedPayload]()).
	Decoder(Document_REPLACED, domain.JSONDecoder[DocumentReplacedPayload]()).
	Decoder(DocumentSignature_REQUESTED, domain.JSONDecoder[DocumentSignatureRequestedPayload]()).
	Decoder(DocumentSignature_SIGNED, domain.JSONDecoder[DocumentSignatureSignedPayload]()).
	Decoder(DocumentSignature_FAILED, domain.JSONDecoder[DocumentSignatureFailedPayload]())
