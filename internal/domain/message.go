package domain

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	LeftSide  uuid.UUID
	ContactId uuid.UUID
	RightSide uuid.UUID
	Text      string
	Sender    bool
	SentAt    time.Time
}
