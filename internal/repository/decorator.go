package repository

import (
	"github.com/google/uuid"
	"tg-p2p-service/internal/domain"
	"time"
)

type Decorator struct{}

func (d Decorator) NewContact(contact domain.Contact) error {
	panic("implement me")
}

func (d Decorator) GetContacts(uuid uuid.UUID) ([]domain.Contact, error) {
	panic("implement me")
}

func (d Decorator) RecordMessage(message domain.Message) error {
	panic("implement me")
}

func (d Decorator) GetMessages(from time.Time, to time.Time, leftSide uuid.UUID, contactId uuid.UUID) ([]domain.Message, error) {
	panic("implement me")
}

func (d Decorator) UpdateMessage(message domain.Message) error {
	panic("implement me")
}

func (d Decorator) GetMessage(message domain.Message) (domain.Message, error) {
	panic("implement me")
}

func (d Decorator) OneWayMessageDelete(message domain.Message) error {
	panic("implement me")
}

func (d Decorator) TwoWayMessageDelete(message domain.Message) error {
	panic("implement me")
}
