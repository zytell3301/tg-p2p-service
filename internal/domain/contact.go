package domain

import "github.com/google/uuid"

type ContactSides struct {
	LeftSide  uuid.UUID
	RightSide uuid.UUID
}

type Contact struct {
	ContactId    uuid.UUID
	ContactSides ContactSides
}
