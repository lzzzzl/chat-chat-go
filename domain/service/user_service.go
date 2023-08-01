package service

import (
	"github/lzzzzl/chat-chat-go/domain/model"
	"github/lzzzzl/chat-chat-go/domain/repository"
)

// UserService encapsulates use case logic for user domain.
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService with the given UserRepository.
func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

// Register creates a new user with the provided username and password.
// It will return an error if there is any problem during the user creation or saving process.
func (s *UserService) Register(username, password string) (*model.User, error) {
	user, err := model.NewUser(username, password)
	if err != nil {
		return nil, err
	}

	// Save the user
	err = s.repo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login verifies the provided username and password and returns the corresponding user.
// It will return an error if the user is not found or the password does not match.
func (s *UserService) Login(username, password string) (*model.User, error) {
	// Get the user
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	// Compare the hashed password
	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}

	return user, nil
}
