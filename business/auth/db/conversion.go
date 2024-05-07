package db

import (
	"time"

	"encore.app/business/auth"
	"encore.app/business/auth/db/postgres"
)

func convertToken(t postgres.Token) auth.Token {
	var validUntil *time.Time
	if t.ValidUntil.Valid {
		validUntil = &t.ValidUntil.Time
	}
	return auth.Token{
		ID:          t.ID,
		UserID:      t.UserID,
		Description: t.Description,
		TokenHash:   t.TokenHash,
		ValidUntil:  validUntil,
		UpdatedAt:   t.UpdatedAt.Time,
		CreatedAt:   t.CreatedAt.Time,
	}
}
