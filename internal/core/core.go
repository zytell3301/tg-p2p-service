package core

import (
	"github.com/google/uuid"
	"github.com/zytell3301/tg-globals/errors"
	uuid_generator "github.com/zytell3301/uuid-generator"
	"tg-p2p-service/internal/domain"
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
	id, err := s.uuidGenerator.GenerateV4()
	contact.ContactId = id
	switch err != nil {
	case true:
		// @TODO report error to central error recorder
		return errors.InternalError{}
	}
	err = s.repository.NewContact(contact)
	switch err != nil {
	case true:
		return errors.InternalError{}
	}
	/*
	 * We flip the contact sides and again insert it into database.
	 */
	err = s.repository.NewContact(domain.Contact{
		ContactId: id,
		ContactSides: domain.ContactSides{
			LeftSide:  contact.ContactSides.RightSide,
			RightSide: contact.ContactSides.LeftSide,
		},
	})
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
