package repositoryController

import (
	"github.com/google/uuid"
	"tg-p2p-service/internal/domain"
	"time"
)

type Repository interface {
	AddContact(contact domain.Contact) (AddContactBatch, error)
	GetContacts(userId uuid.UUID) ([]domain.Contact, error)
	RecordMessage(message domain.Message) (RecordMessageBatch, error)
	GetMessages(from time.Time, to time.Time, leftSide uuid.UUID, contactId uuid.UUID) ([]domain.Message, error)
}

type AddContactBatch interface {
	// AddContactToBatch only adds a query to batch and won't execute the batch.
	// You must explicitly call ExecuteOperation method to execute the batch.
	// If batch feature is not available, you can separately run the queries
	// but you will lose the database consistency on failure on any operation.
	AddContactToBatch(contact domain.Contact) error

	// After adding queries to batch, you must call ExecuteOperation method
	// or the changes will be lost
	ExecuteOperation() error
}

type RecordMessageBatch interface {
	// AddMessageToBatch only adds a query to batch and won't execute the batch.
	// You must explicitly call ExecuteOperation method to execute the batch.
	// If batch feature is not available, you can separately run the queries
	// but you will lose the database consistency on failure on any operation.
	AddMessageToBatch(message domain.Message) error

	// After adding queries to batch, you must call ExecuteOperation method
	// or the changes will be lost
	ExecuteOperation() error
}
