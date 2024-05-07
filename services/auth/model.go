package auth

import "encore.dev/types/uuid"

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password" encore:"sensitive"`
}

type Token struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
}

type PlainToken struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Token       string    `json:"token"`
}
