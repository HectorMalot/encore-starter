// Package appdb declares the application database, contains the SQL for
// database migrations and seeding.
package appdb

import (
	"encore.dev/storage/sqldb"
)

// This represents the database for this system. Encore will create and
// manage this database for us. The name has to be a literal string.
var _ = sqldb.NewDatabase("app", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
