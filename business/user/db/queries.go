// Package db contains the storage logic for users.
//
// SQL queries are defined in `queries.sql` and `sqlc` is used to
// generate the relevant Go code in the `postgres` package.
package db

import (
	"context"

	"encore.app/business/user"
	"encore.app/business/user/db/postgres"
	"encore.dev/types/uuid"
)

func (r *Repository) CreateUser(ctx context.Context, newUser user.User) (user.User, error) {
	u, err := r.queries.CreateUser(ctx, postgres.CreateUserParams{
		Email:        newUser.Email,
		Roles:        newUser.Roles,
		PasswordHash: newUser.PasswordHash,
	})
	if err != nil {
		return user.User{}, err
	}
	return convertUser(u), nil
}

func (r *Repository) ListUsers(ctx context.Context) ([]user.User, error) {
	users, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	return convertUsers(users), nil
}

func (r *Repository) GetUserByID(ctx context.Context, userID uuid.UUID) (user.User, error) {
	u, err := r.queries.GetUserByID(ctx, userID)
	if err != nil {
		return user.User{}, err
	}
	return convertUser(u), nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	u, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return user.User{}, err
	}
	return convertUser(u), nil
}

func (r *Repository) UpdateUser(ctx context.Context, userID uuid.UUID, updatedUser user.User) (user.User, error) {
	u, err := r.queries.UpdateUser(ctx, postgres.UpdateUserParams{
		ID:           userID,
		Email:        updatedUser.Email,
		Roles:        updatedUser.Roles,
		PasswordHash: updatedUser.PasswordHash,
	})
	if err != nil {
		return user.User{}, err
	}
	return convertUser(u), nil
}

func (r *Repository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	err := r.queries.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
