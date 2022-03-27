package core

import "tg-p2p-service/internal/domain"

// Message logger logs any message changes such as deleting or editing a message.
type MessageLogger interface {
	LogMessage(message domain.Message)
}
