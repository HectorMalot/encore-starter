package auth

import (
	"time"

	"encore.dev/types/uuid"
)

type Token struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Description string
	TokenHash   string
	ValidUntil  *time.Time
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type NewToken struct {
	UserID      uuid.UUID
	Description string
	ValidUntil  *time.Time
}

type PlainToken struct {
	Token
	PlainToken string
}
