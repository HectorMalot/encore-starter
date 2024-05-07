package auth

import (
	"context"
	"fmt"
	"time"

	"encore.app/business/auth"
)

const loginTokenValidity = time.Hour * 24

// Login validates a user's credentials and returns a token if they are valid.
//
//encore:api public path=/v1/auth/login method=POST
func (s *Service) Login(ctx context.Context, creds LoginCredentials) (*PlainToken, error) {
	u, err := s.user.ValidateCredentials(ctx, creds.Email, creds.Password)
	if err != nil {
		return nil, err
	}

	lifetime := time.Now().Add(loginTokenValidity)
	t, err := s.auth.Create(ctx, auth.NewToken{
		UserID:      u.ID,
		Description: fmt.Sprintf("Login token: %s", time.Now().String()),
		ValidUntil:  &lifetime,
	})
	if err != nil {
		return nil, err
	}

	return &PlainToken{
		ID:          t.ID,
		Description: t.Description,
		Token:       t.PlainToken,
	}, nil
}
