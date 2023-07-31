package model

import "testing"

func TestNewRoom(t *testing.T) {
	name := "Room 1"

	room, err := NewRoom(name)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if room.Name != name {
		t.Errorf("Expected Name to be %s, got %s", name, room.Name)
	}
}
