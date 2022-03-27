package repositoryController

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	ErrorReporter "github.com/zytell3301/tg-error-reporter"
	errors2 "github.com/zytell3301/tg-globals/errors"
	"tg-p2p-service/internal/domain"
	"time"
)

const (
	newContactErrorMessage      = "an error occurred while adding a new contact. Error message: %s"
	recordMessageErrorMessage   = "an error occurred while recording a message. Error message: %s"
	getMessageRangeErrorMessage = "an error occurred while getting a message range. Error message: %s"
	getMessageErrorMessage      = "an error occurred while getting a message. Error message: %s"
	oneWayMessgeDelete          = "an error occurred while one way deleting a message. Error message: %s"
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
	messages, err := d.repository.GetMessages(from, to, leftSide, contactId)
	switch err != nil {
	case true:
		switch errors.As(err, &errors2.EntityNotFound{}) {
		case true:
			// In this case the error is EntityNotFound. So we let the caller
			// find out that query resulted in empty response
			return nil, err
		default:
			d.reportError(getMessageRangeErrorMessage, err.Error())
			return nil, errors2.InternalError{}
		}
	}
	return messages, nil
}

func (d Decorator) UpdateMessage(message domain.Message) error {
	batch, err := d.repository.UpdateMessage(message)
	switch err != nil {
	case true:
		return err
	}
	err = batch.AddUpdateToBatch(domain.Message{
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
		return err
	}
	err = batch.ExecuteOperation()
	switch err != nil {
	case true:
		return err
	}

	return nil
}

func (d Decorator) GetMessage(message domain.Message) (domain.Message, error) {
	message, err := d.repository.GetMessage(message)
	switch err != nil {
	case true:
		switch errors.As(err, &errors2.EntityNotFound{}) {
		case true:
			// Let the caller know that query resulted in empty response
			return domain.Message{}, err
		default:
			d.reportError(getMessageErrorMessage, err.Error())
			return domain.Message{}, err
		}
	}
	return message, nil
}

func (d Decorator) OneWayMessageDelete(message domain.Message) error {
	batch, err := d.oneWayMessageDelete(message)
	switch err != nil {
	case true:
		d.reportError(oneWayMessgeDelete, err.Error())
		return errors2.InternalError{}
	}
	err = batch.ExecuteOperation()
	switch err != nil {
	case true:
		d.reportError(oneWayMessgeDelete, err.Error())
		return errors2.InternalError{}
	}

	return nil
}

func (d Decorator) oneWayMessageDelete(message domain.Message) (DeleteMessageBatch, error) {
	batch, err := d.repository.DeleteMessage(message)
	return batch, err
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
