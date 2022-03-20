package core

import (
	"github.com/google/uuid"
	"tg-p2p-service/internal/domain"
)

type Repository interface {
	NewContact(contact domain.Contact) error
	GetContacts(uuid uuid.UUID) ([]domain.Contact, error)
}
