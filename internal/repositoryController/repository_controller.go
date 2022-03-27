package repositoryController

import (
	"fmt"
	"github.com/google/uuid"
	ErrorReporter "github.com/zytell3301/tg-error-reporter"
	"tg-p2p-service/internal/domain"
	"time"
)

const (
	newContactErrorMessage    = "an error occurred while adding a new contact. Error message: %s"
	recordMessageErrorMessage = "an error occurred while recording a message. Error message: %s"
)

type Decorator struct {
	repository  Repository
	reporter    ErrorReporter.Reporter
	serviceInfo domain.ServiceInfo
}

func (d Decorator) NewContact(contact domain.Contact) error {
	batch, err := d.repository.AddContact(contact)
	switch err != nil {
	case true:
		d.reportError(newContactErrorMessage, err.Error())
	}

	// Swap contact sides
	err = batch.AddContactToBatch(domain.Contact{
		ContactId: contact.ContactId,
		ContactSides: domain.ContactSides{
			LeftSide:  contact.ContactSides.RightSide,
			RightSide: contact.ContactSides.LeftSide,
		},
	})
	switch err != nil {
	case true:
		d.reportError(newContactErrorMessage, err.Error())
	}

	err = batch.ExecuteOperation()
	switch err != nil {
	case true:
		d.reportError(newContactErrorMessage, err.Error())
		return err
	}

	return nil
}

func (d Decorator) GetContacts(userId uuid.UUID) ([]domain.Contact, error) {
	return d.repository.GetContacts(userId)
}

func (d Decorator) RecordMessage(message domain.Message) error {
	batch, err := d.repository.RecordMessage(message)
	switch err != nil {
	case true:
		d.reportError(recordMessageErrorMessage, err.Error())
		return err
	}

	// Swap message sides and reinsert it into batch
	err = batch.AddMessageToBatch(domain.Message{
		LeftSide:  message.RightSide,
		ContactId: message.ContactId,
		RightSide: message.LeftSide,
		Text:      message.Text,
		Sender:    flipSender(message.Sender),
		SentAt:    message.SentAt,
		MessageId: message.MessageId,
	})
	switch err != nil {
	case true:
		d.reportError(recordMessageErrorMessage, err.Error())
		return err
	}

	err = batch.ExecuteOperation()
	switch err != nil {
	case true:
		return err
	}

	return nil
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

// template obeys fmt.Sprintf rules and params will be replaced with placeholders
func (d Decorator) reportError(template string, params ...string) {
	d.reporter.Report(ErrorReporter.Error{
		ServiceGroupId: d.serviceInfo.ServiceGroupId,
		InstanceId:     d.serviceInfo.InstanceId,
		Message:        fmt.Sprintf(template, params),
	})
}

func flipSender(sender bool) bool {
	switch sender {
	case true:
		return false
	default:
		return true
	}
}
