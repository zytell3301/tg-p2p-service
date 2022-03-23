package core

import (
	"github.com/google/uuid"
	"github.com/zytell3301/tg-globals/errors"
	uuid_generator "github.com/zytell3301/uuid-generator"
	"tg-p2p-service/internal/domain"
	"time"
)

type Service struct {
	repository    Repository
	uuidGenerator uuid_generator.Generator
}

type ServiceConfig struct {
}

type Dependencies struct {
	Repository    Repository
	UuidGenerator uuid_generator.Generator
}

func NewMessagesCore(config ServiceConfig, dependencies Dependencies) Service {
	return Service{
		repository:    dependencies.Repository,
		uuidGenerator: dependencies.UuidGenerator,
	}
}

func (s Service) NewContact(contact domain.Contact) error {
	id := s.uuidGenerator.GenerateV4()
	contact.ContactId = id
	err := s.repository.NewContact(contact)
	switch err != nil {
	case true:
		return errors.InternalError{}
	}

	return nil
}

func (s Service) GetContacts(userId uuid.UUID) ([]domain.Contact, error) {
	contacts, err := s.repository.GetContacts(userId)
	switch err != nil {
	case true:
		return []domain.Contact{}, errors.InternalError{}
	}
	return contacts, nil
}

func (s Service) SendMessage(message domain.Message) error {
	message.SentAt = time.Now()
	err := s.repository.RecordMessage(message)
	switch err != nil {
	case true:
		return errors.InternalError{}
	}

	return nil
}

func (s Service) GetMessages(from time.Time, to time.Time, leftSide uuid.UUID, contactId uuid.UUID) ([]domain.Message, error) {
	messages, err := s.repository.GetMessages(from, to, leftSide, contactId)
	switch err != nil {
	case true:
		return []domain.Message{}, errors.InternalError{}
	}

	return messages, nil
}
