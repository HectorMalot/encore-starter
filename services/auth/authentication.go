package auth

import (
	"context"
	"errors"

	"encore.app/business/user"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
)

var (
	ErrInvalidCurrentUser = errors.New("current user not valid for this operation")
)

type Data struct {
	UserID uuid.UUID
	User   *user.User
}

// AuthHandler handled authentication for all endpoints defined with 'auth' in the 'encore:api' directive.
// The token is passed in the Authorization header as a Bearer token.
//
//encore:authhandler
func (s *Service) AuthHandler(ctx context.Context, token string) (auth.UID, *Data, error) {
	userID, err := s.auth.Validate(ctx, token)
	if err != nil {
		return "", nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "invalid token",
		}
	}

	// Fetch the user from the database
	u, err := s.user.Get(ctx, userID)
	if err != nil {
		return "", nil, &errs.Error{
			Code:    errs.Internal,
			Message: "invalid user",
		}
	}

	return auth.UID(userID.String()), &Data{UserID: userID, User: &u}, nil
}

// Returns the current user from the authentication system
func CurrentUser() (*user.User, error) {
	data, ok := auth.Data().(*Data)
	if !ok || data == nil {
		return nil, ErrInvalidCurrentUser
	}
	return data.User, nil
}
