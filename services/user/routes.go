package user

import (
	"context"

	"encore.app/business/user"
	"encore.app/utils/slices"
	"encore.dev/types/uuid"
)

// CreateUser creates a new user
//
//encore:api auth path=/v1/users method=POST tag:authorize
func (s *Service) CreateUser(ctx context.Context, u NewUser) (*User, error) {
	nu, err := s.user.Create(ctx, user.NewUser{
		Email:    u.Email,
		Roles:    u.Roles,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    nu.ID,
		Email: nu.Email,
		Roles: nu.Roles,
	}, nil
}

// GetUser retrieves a user by ID
//
//encore:api auth path=/v1/users/:id method=GET tag:authorize
func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
	u, err := s.user.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    u.ID,
		Email: u.Email,
		Roles: u.Roles,
	}, nil
}

// ListUsers retrieves all users
//
//encore:api auth path=/v1/users method=GET tag:authorize
func (s *Service) ListUsers(ctx context.Context) (*List, error) {
	users, err := s.user.List(ctx)
	if err != nil {
		return nil, err
	}

	us := slices.Map(users, func(u user.User) *User {
		return &User{
			ID:    u.ID,
			Email: u.Email,
			Roles: u.Roles,
		}
	})

	return &List{us}, nil
}

// UpdateUser updates a user by ID
//
//encore:api auth path=/v1/users/:id method=PUT tag:authorize
func (s *Service) UpdateUser(ctx context.Context, id uuid.UUID, u NewUser) (*User, error) {
	nu, err := s.user.Update(ctx, id, user.NewUser{
		Email:    u.Email,
		Roles:    u.Roles,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    nu.ID,
		Email: nu.Email,
		Roles: nu.Roles,
	}, nil
}

// DeleteUser deletes a user by ID
//
//encore:api auth path=/v1/users/:id method=DELETE tag:authorize
func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.user.Delete(ctx, id)
}
