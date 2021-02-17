package domain

import "github.com/lindseypoche/SELU_ACM/api/internal/utils/errors/rest"

// UserRepository is an interface for blog repositories
type UserRepository interface {
	Save(user *User) (*User, rest.Err)
	GetByID(id int64) (*User, rest.Err)
}
