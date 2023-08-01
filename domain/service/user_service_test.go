package service

import (
	"github/lzzzzl/chat-chat-go/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// UserRepositoryMock is a mock for UserRepository
type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Save(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) DeleteByID(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindByID(id int) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestUserService(t *testing.T) {
	// Create a mock UserRepository
	mockRepo := new(UserRepositoryMock)

	// Create a new UserService with the mock UserRepository
	service := NewUserService(mockRepo)

	// Devine the behavior for the mock UserRepository's Save method
	mockRepo.On("Save", mock.Anything).Return(nil)

	// Test the Register method
	user, err := service.Register("testuser", "testpassword")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	// Define the behavior for the mock UserRepository's FindByUsername method
	mockRepo.On("FindByUsername", "testuser").Return(user, nil)

	user, err = service.Login("testuser", "testpassword")
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
