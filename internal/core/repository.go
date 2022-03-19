package core

import "tg-p2p-service/internal/domain"

type Repository interface {
	NewContact(contact domain.Contact) error
}
