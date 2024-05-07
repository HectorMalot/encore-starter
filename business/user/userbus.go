package user

import (
	"context"

	"encore.dev/rlog"
	"encore.dev/types/uuid"
)

type Storer interface {
	CreateUser(ctx context.Context, newUser User) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, updatedUser User) (User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

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

// Create generates a new User and saves it.
func (b *Business) Create(ctx context.Context, newUser NewUser) (User, error) {
	hash, err := hashPassword(newUser.Password)
	if err != nil {
		return User{}, err
	}

	return b.storer.CreateUser(ctx, User{
		Email:        newUser.Email,
		Roles:        newUser.Roles,
		PasswordHash: hash,
	})
}

// Get retrieves a User by its ID.
func (b *Business) Get(ctx context.Context, id uuid.UUID) (User, error) {
	return b.storer.GetUserByID(ctx, id)
}

// List retrieves all Users.
func (b *Business) List(ctx context.Context) ([]User, error) {
	return b.storer.ListUsers(ctx)
}

// Update modifies a User by its ID.
func (b *Business) Update(ctx context.Context, id uuid.UUID, updatedUser NewUser) (User, error) {
	u, err := b.storer.GetUserByID(ctx, id)
	if err != nil {
		return User{}, err
	}

	if updatedUser.Password != "" {
		hash, err := hashPassword(updatedUser.Password)
		if err != nil {
			return User{}, err
		}
		u.PasswordHash = hash
	}

	u.Email = updatedUser.Email
	u.Roles = updatedUser.Roles

	return b.storer.UpdateUser(ctx, id, u)
}

// Delete removes a User by its ID.
func (b *Business) Delete(ctx context.Context, id uuid.UUID) error {
	return b.storer.DeleteUser(ctx, id)
}

// ValidateCredentials checks if the provided email and password match a User.
func (b *Business) ValidateCredentials(ctx context.Context, email, password string) (User, error) {
	u, err := b.storer.GetUserByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	err = comparePassword(u.PasswordHash, password)
	if err != nil {
		return User{}, err
	}

	return u, nil
}
