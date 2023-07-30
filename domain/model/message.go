package model

import (
	"errors"
	"time"
)

// Message represents a chat message.
type Message struct {
	ID        int
	UserID    int
	RoomID    int
	Content   string
	CreatedAt time.Time
}

// NewMessage createds a new Message
// This function will also validate the input.
func NewMessage(userID int, roomID int, content string, createdAt time.Time) (*Message, error) {
	if userID <= 0 {
		return nil, errors.New("userID must not be empty")
	}
	if roomID <= 0 {
		return nil, errors.New("roomID must not be empty")
	}
	if content == "" {
		return nil, errors.New("content must not be empty")
	}

	return &Message{
		UserID:    userID,
		RoomID:    roomID,
		Content:   content,
		CreatedAt: createdAt,
	}, nil
}
