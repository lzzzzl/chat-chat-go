package model

import (
	"errors"
	"time"
)

type Message struct {
	ID        string
	SenderID  string
	RoomID    string
	Content   string
	CreatedAt time.Time
}

func NewMessage(id string, senderID string, roomID string, content string, createdAt time.Time) (*Message, error) {
	if id == "" {
		return nil, errors.New("id must not be empty")
	}
	if senderID == "" {
		return nil, errors.New("senderID must not be empty")
	}
	if roomID == "" {
		return nil, errors.New("roomID must not be empty")
	}
	if content == "" {
		return nil, errors.New("content must not be empty")
	}

	return &Message{
		ID:        id,
		SenderID:  senderID,
		RoomID:    roomID,
		Content:   content,
		CreatedAt: createdAt,
	}, nil
}
