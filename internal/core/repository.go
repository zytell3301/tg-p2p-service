package core

import (
	"github.com/google/uuid"
	"tg-p2p-service/internal/domain"
)

type Repository interface {
	NewContact(contact domain.Contact) error // @TODO The contact must be flipped and reinserted into database
	GetContacts(uuid uuid.UUID) ([]domain.Contact, error)
	RecordMessage(message domain.Message) error //@TODO flip message sides and reinsert into database
}
