package service

import (
	"github/lzzzzl/chat-chat-go/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// RoomRepositoryMock is a mock implementation of the RoomRepository.
type RoomRepositoryMock struct {
	mock.Mock
}

// Save is a mock method that simulates the Save method of the RoomRepository.
func (r *RoomRepositoryMock) Save(room *model.Room) error {
	args := r.Called(room)
	return args.Error(0)
}

// FindAll is a mock method that simulates the FindAll method of the RoomRepository.
func (r *RoomRepositoryMock) FindAll() ([]*model.Room, error) {
	args := r.Called()
	return args.Get(0).([]*model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) FindByID(id int) (*model.Room, error) {
	args := r.Called(0)
	return args.Get(0).(*model.Room), args.Error(1)
}

func TestRoomService(t *testing.T) {
	// Create a mock RoomRepository
	mockRepo := new(RoomRepositoryMock)

	// Create a new RoomService with the mock RoomRepository
	service := NewRoomService(mockRepo)

	// Define the behavior for the mock RoomRepository's Save method
	room := &model.Room{Name: "Test Room"}
	mockRepo.On("Save", room).Return(nil)

	// Test the Create method
	createdRoom, err := service.Create(room.Name)
	assert.NoError(t, err)
	assert.NotNil(t, createdRoom)

	// Define the behavior for the mock RoomRepository's FindAll method
	rooms := []*model.Room{room}
	mockRepo.On("FindAll").Return(rooms, nil)

	// Test the GetAll method
	returnedRooms, err := service.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, returnedRooms)
}
