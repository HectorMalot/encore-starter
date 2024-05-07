package db

import (
	"encore.app/business/auth"
	"encore.app/business/auth/db/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Ensure Repository implements the storer interface
var _ auth.Storer = &Repository{}

type Repository struct {
	queries *postgres.Queries
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{postgres.New(pool)}
}
