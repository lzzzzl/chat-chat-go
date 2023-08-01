package service

import (
	"github/lzzzzl/chat-chat-go/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MessageRepositoryMock struct {
	mock.Mock
}

func (m *MessageRepositoryMock) FindByID(id int) (*model.Message, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Message), args.Error(1)
}

func (m *MessageRepositoryMock) Save(message *model.Message) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *MessageRepositoryMock) DeleteByID(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MessageRepositoryMock) FindMessagesByRoomID(roomID int) ([]*model.Message, error) {
	args := m.Called(roomID)
	return args.Get(0).([]*model.Message), args.Error(1)
}

func TestMessageService(t *testing.T) {
	// Create a mock MessageRepository
	mockRepo := new(MessageRepositoryMock)

	// Create a new MessageService with the mock MessageRepository
	service := NewMessageService(mockRepo)

	// Devine the behavior for the mock MessageRepository's Save method
	mockRepo.On("Save", mock.Anything).Return(nil)

	// Test the Save method
	message, err := service.Save(1, 1, "Test Content")
	assert.NoError(t, err)
	assert.NotNil(t, message)

	// Define the behavior for the mock MessageRepository's FindMessagesByRoomID method
	messages := []*model.Message{message}
	mockRepo.On("FindMessagesByRoomID", mock.Anything).Return(messages, nil)

	messages, err = service.FindMessagesByRoomID(1)
	assert.NoError(t, err)
	assert.NotNil(t, messages)
}
