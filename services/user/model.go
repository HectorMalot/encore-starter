package user

import (
	"fmt"
	"regexp"

	"encore.dev/types/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Roles []string  `json:"roles"`
}

type NewUser struct {
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
	Password string   `json:"password"`
}

// Encore will call this method to validate the input
func (nu NewUser) Validate() error {
	// Validate the email with a regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(nu.Email) {
		return fmt.Errorf("invalid email address")
	}

	// Validate the password length
	if len(nu.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	return nil
}

type List struct {
	Users []*User `json:"users"`
}
