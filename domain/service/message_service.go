package service

import (
	"github/lzzzzl/chat-chat-go/domain/model"
	"github/lzzzzl/chat-chat-go/domain/repository"
	"time"
)

// MessageService provides message-related operations.
type MessageService struct {
	repo repository.MessageRepository
}

// NewMessageService returns a new MessageService.
func NewMessageService(r repository.MessageRepository) *MessageService {
	return &MessageService{repo: r}
}

// Save saves a new message to the repository.
func (s *MessageService) Save(userID, roomID int, content string) (*model.Message, error) {
	message, err := model.NewMessage(userID, roomID, content, time.Now())
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (s *MessageService) FindMessagesByRoomID(roomID int) ([]*model.Message, error) {
	return s.repo.FindMessagesByRoomID(roomID)
}
