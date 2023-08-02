package service

import (
	"github/lzzzzl/chat-chat-go/domain/model"
	"github/lzzzzl/chat-chat-go/domain/repository"
)

// RoomService represents the application service for managing chat rooms.
type RoomService struct {
	repo repository.RoomRepository
}

// NewRoomService creates a new RoomService with the given RoomRepository.
func NewRoomService(r repository.RoomRepository) *RoomService {
	return &RoomService{repo: r}
}

// Create creates a new chat room with the given name.
func (s *RoomService) Create(name string) (*model.Room, error) {
	room := &model.Room{Name: name}

	// Save the room
	err := s.repo.Save(room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// GetAll retrieves all chat rooms.
func (s *RoomService) GetAll() ([]*model.Room, error) {
	return s.repo.FindAll()
}
