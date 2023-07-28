package model

import "testing"

func TestNewRoom(t *testing.T) {
	id := "1"
	name := "Room 1"

	room, err := NewRoom(id, name)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if room.ID != id {
		t.Fatalf("Expected ID to be %s, got %s", id, room.ID)
	}
	if room.Name != name {
		t.Errorf("Expected Name to be %s, got %s", name, room.Name)
	}
}
