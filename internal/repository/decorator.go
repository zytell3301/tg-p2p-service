package repository

import (
	"github.com/google/uuid"
	"tg-p2p-service/internal/domain"
	"time"
)

type Decorator struct{}

func (r Decorator) NewContact(contact domain.Contact) error {
	panic("implement me")
}

func (r Decorator) GetContacts(uuid uuid.UUID) ([]domain.Contact, error) {
	panic("implement me")
}

func (r Decorator) RecordMessage(message domain.Message) error {
	panic("implement me")
}

func (r Decorator) GetMessages(from time.Time, to time.Time, leftSide uuid.UUID, contactId uuid.UUID) ([]domain.Message, error) {
	panic("implement me")
}

func (r Decorator) UpdateMessage(message domain.Message) error {
	panic("implement me")
}

func (r Decorator) GetMessage(message domain.Message) (domain.Message, error) {
	panic("implement me")
}

func (r Decorator) OneWayMessageDelete(message domain.Message) error {
	panic("implement me")
}

func (r Decorator) TwoWayMessageDelete(message domain.Message) error {
	panic("implement me")
}
