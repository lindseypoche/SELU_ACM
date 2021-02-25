package domain

import (
	"github.com/lindseypoche/SELU_ACM/api/internal/utils/crypto_utils"
	"github.com/lindseypoche/SELU_ACM/api/internal/utils/date_utils"
	"github.com/lindseypoche/SELU_ACM/api/internal/utils/errors/rest"
)

type userService struct {
	userRepo UserRepository
}

// UserService is an interface for users
type UserService interface {
	CreateUser(*User) (*User, rest.Err)
	GetUser(int64) (*User, rest.Err)
}

// NewUserService creates a new user service
func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser creates the user data that is received from the CreateUser controller
func (s *userService) CreateUser(user *User) (*User, rest.Err) {

	// validate user data
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMD5(user.Password)

	result, err := s.userRepo.Save(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUser gets and returns a user by their user id
func (s *userService) GetUser(userID int64) (*User, rest.Err) {

	// create new instance of user and give it userID
	// user := &User{ID: userID}

	// check if user id exists in db
	result, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
