package core

import (
	"github.com/google/uuid"
	"tg-p2p-service/internal/domain"
	"time"
)

type Repository interface {
	NewContact(contact domain.Contact) error // @TODO The contact must be flipped and reinserted into database
	GetContacts(userId uuid.UUID) ([]domain.Contact, error)
	RecordMessage(message domain.Message) error //@TODO flip message sides and reinsert into database
	GetMessages(from time.Time, to time.Time, leftSide uuid.UUID, contactId uuid.UUID) ([]domain.Message, error)
	UpdateMessage(message domain.Message) error // @TODO make sure that both message replications will get updated

	// Since we can't trust to user provided information, must fetch the message again from database
	GetMessage(message domain.Message) (domain.Message, error)

	OneWayMessageDelete(message domain.Message) error // @TODO handle logging operation in repository decorator
	TwoWayMessageDelete(message domain.Message) error // @TODO handle logging operation in repository decorator
}
