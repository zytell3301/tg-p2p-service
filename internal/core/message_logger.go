package core

import "tg-p2p-service/internal/domain"

// Message logger logs any message changes such as deleting or editing a message.
type MessageLogger interface {
	LogOneWayDeletedMessage(message domain.Message) error
	LogTwoWayDeletedMessage(message domain.Message) error
	LogEditedMessage(message domain.Message) error
}
