package user

import "encore.dev/types/uuid"

type User struct {
	ID           uuid.UUID
	Email        string
	Roles        []string
	PasswordHash []byte
}

type NewUser struct {
	Email    string
	Roles    []string
	Password string
}
