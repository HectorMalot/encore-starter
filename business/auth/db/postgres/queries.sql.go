// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package postgres

import (
	"context"

	"encore.dev/types/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const deleteToken = `-- name: DeleteToken :exec
DELETE FROM tokens WHERE id = $1
`

func (q *Queries) DeleteToken(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteToken, id)
	return err
}

const getToken = `-- name: GetToken :one
SELECT id, user_id, description, token_hash, valid_until, updated_at, created_at FROM tokens WHERE token_hash = $1
`

func (q *Queries) GetToken(ctx context.Context, tokenHash string) (Token, error) {
	row := q.db.QueryRow(ctx, getToken, tokenHash)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Description,
		&i.TokenHash,
		&i.ValidUntil,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listTokens = `-- name: ListTokens :many
SELECT id, user_id, description, token_hash, valid_until, updated_at, created_at FROM tokens WHERE user_id = $1
`

func (q *Queries) ListTokens(ctx context.Context, userID uuid.UUID) ([]Token, error) {
	rows, err := q.db.Query(ctx, listTokens, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Description,
			&i.TokenHash,
			&i.ValidUntil,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const storeToken = `-- name: StoreToken :one
INSERT INTO tokens (user_id, description, token_hash, valid_until) VALUES ($1, $2, $3, $4) RETURNING id
`

type StoreTokenParams struct {
	UserID      uuid.UUID
	Description string
	TokenHash   string
	ValidUntil  pgtype.Timestamptz
}

func (q *Queries) StoreToken(ctx context.Context, arg StoreTokenParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, storeToken,
		arg.UserID,
		arg.Description,
		arg.TokenHash,
		arg.ValidUntil,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}