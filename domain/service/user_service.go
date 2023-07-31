package service

import (
	"github/lzzzzl/chat-chat-go/domain/model"
	"github/lzzzzl/chat-chat-go/domain/repository"
)

// UserService ...
type UserService struct {
	repo repository.UserRepository
}

// NewUserService ...
func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

// Register ...
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

// Login ...
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
