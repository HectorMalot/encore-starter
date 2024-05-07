package db

import (
	"encore.app/business/user"
	"encore.app/business/user/db/postgres"
	"encore.app/utils/slices"
)

func convertUser(u postgres.User) user.User {
	return user.User{
		ID:           u.ID,
		Email:        u.Email,
		Roles:        u.Roles,
		PasswordHash: u.PasswordHash,
	}
}

func convertUsers(users []postgres.User) []user.User {
	return slices.Map(users, convertUser)
}
