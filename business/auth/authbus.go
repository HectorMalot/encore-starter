// Package auth implements business logic for user tokens, typically
// used for authentication and authorization.
//
// Tokens are hashed (SHA-256) and encoded (base64) before being stored.
package auth

import (
	"context"
	"fmt"
	"time"

	"encore.dev/rlog"
	"encore.dev/types/uuid"
)

type Storer interface {
	GetToken(ctx context.Context, token_hash string) (Token, error)
	StoreToken(ctx context.Context, token Token) error
	ListTokens(ctx context.Context, userID uuid.UUID) ([]Token, error)
	DeleteToken(ctx context.Context, id uuid.UUID) error
}

const TokenLength = 32 // 256 bits

var (
	ErrTokenExpired = fmt.Errorf("token expired")
)

// Business manages the set of APIs for authentication tokens.
type Business struct {
	log    rlog.Ctx
	storer Storer
}

// NewBusiness constructs an authentication business API for use.
func NewBusiness(log rlog.Ctx, storer Storer) *Business {
	return &Business{
		log:    log,
		storer: storer,
	}
}

// Create generates a new cryptographically secure token and saves it.
func (b *Business) Create(ctx context.Context, np NewToken) (PlainToken, error) {
	plainToken, err := b.generateToken(TokenLength)
	if err != nil {
		b.log.Error("failed to generate token", "error", err)
		return PlainToken{}, err
	}

	token := Token{
		UserID:      np.UserID,
		Description: np.Description,
		TokenHash:   b.hash(plainToken),
		ValidUntil:  np.ValidUntil,
	}

	if err := b.storer.StoreToken(ctx, token); err != nil {
		b.log.Error("failed to store token", "error", err)
		return PlainToken{}, err
	}

	return PlainToken{
		Token:      token,
		PlainToken: plainToken,
	}, nil
}

// Validate checks if a token is valid and returns an error if the token is invalid or expired.
func (b *Business) Validate(ctx context.Context, t string) (uuid.UUID, error) {
	token, err := b.storer.GetToken(ctx, b.hash(t))
	if err != nil {
		return uuid.Nil, err
	}

	if token.ValidUntil != nil && token.ValidUntil.Before(time.Now()) {
		return uuid.Nil, ErrTokenExpired
	}

	return token.UserID, nil
}
