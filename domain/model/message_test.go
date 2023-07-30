package model

import (
	"testing"
	"time"
)

func TestNewMessage(t *testing.T) {
	userID := 1
	roomID := 1
	content := "Hello, world!"
	createdAt := time.Now()

	message, err := NewMessage(userID, roomID, content, createdAt)

	if err != nil {
		t.Fatalf("Excepted no error, go %v", err)
	}
	if message.UserID != userID {
		t.Errorf("Expected SenderID to be %d, got %d", userID, message.UserID)
	}
	if message.RoomID != roomID {
		t.Errorf("Expected RoomID to be %d, got %d", roomID, message.RoomID)
	}
	if message.Content != content {
		t.Errorf("Expected Content to be %s, got %s", content, message.Content)
	}
	if !message.CreatedAt.Equal(createdAt) {
		t.Errorf("Expected CreatedAt to be %v, got %v", createdAt, message.CreatedAt)
	}
}
