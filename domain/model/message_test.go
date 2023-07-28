package model

import (
	"testing"
	"time"
)

func TestNewMessage(t *testing.T) {
	id := "1"
	senderID := "user1"
	roomID := "room1"
	content := "Hello, world!"
	createdAt := time.Now()

	message, err := NewMessage(id, senderID, roomID, content, createdAt)

	if err != nil {
		t.Fatalf("Excepted no error, go %v", err)
	}
	if message.ID != id {
		t.Errorf("Excepted ID to be %s, got %s", id, message.ID)
	}
	if message.SenderID != senderID {
		t.Errorf("Expected SenderID to be %s, got %s", senderID, message.SenderID)
	}
	if message.RoomID != roomID {
		t.Errorf("Expected RoomID to be %s, got %s", roomID, message.RoomID)
	}
	if message.Content != content {
		t.Errorf("Expected Content to be %s, got %s", content, message.Content)
	}
	if !message.CreatedAt.Equal(createdAt) {
		t.Errorf("Expected CreatedAt to be %v, got %v", createdAt, message.CreatedAt)
	}
}
