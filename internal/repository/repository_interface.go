package repository

import "tg-p2p-service/internal/domain"

type Repository interface {
	AddContact(contact domain.Contact) (error, AddContactBatch)
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