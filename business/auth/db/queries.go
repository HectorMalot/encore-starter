// Package db contains the storage logic for auth tokens
//
// SQL queries are defined in `queries.sql` and `sqlc` is used to
// generate the relevant Go code in the `postgres` package.
package db

import (
	"context"

	"encore.app/business/auth"
	"encore.app/business/auth/db/postgres"
	"encore.app/utils/slices"
	"encore.dev/types/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) GetToken(ctx context.Context, token_hash string) (auth.Token, error) {
	t, err := r.queries.GetToken(ctx, token_hash)
	if err != nil {
		return auth.Token{}, err
	}
	return convertToken(t), nil
}

func (r *Repository) StoreToken(ctx context.Context, token auth.Token) error {
	var validUntil pgtype.Timestamptz
	if token.ValidUntil != nil {
		validUntil.Time = *token.ValidUntil
		validUntil.Valid = true
	}
	_, err := r.queries.StoreToken(ctx, postgres.StoreTokenParams{
		UserID:      token.UserID,
		Description: token.Description,
		TokenHash:   token.TokenHash,
		ValidUntil:  validUntil,
	})
	return err
}

func (r *Repository) ListTokens(ctx context.Context, userID uuid.UUID) ([]auth.Token, error) {
	tokens, err := r.queries.ListTokens(ctx, userID)
	if err != nil {
		return nil, err
	}

	return slices.Map(tokens, convertToken), nil
}

func (r *Repository) DeleteToken(ctx context.Context, id uuid.UUID) error {
	return r.queries.DeleteToken(ctx, id)
}
