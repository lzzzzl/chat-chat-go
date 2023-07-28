package model

import (
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {
	id := "1"
	username := "test"
	password := "password"

	user, err := NewUser(id, username, password)
	fmt.Printf("password, %s", user.PasswordHash)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if user.ID != id {
		t.Errorf("Expected ID to be %s, got %s", id, user.ID)
	}
	if user.Username != username {
		t.Errorf("Expected Username to be %s, got %s", username, user.Username)
	}
	if user.CheckPassword(password) != true {
		t.Errorf("Expected password to match")
	}
}

func TestCheckPassword(t *testing.T) {
	id := "1"
	username := "test"
	password := "password"
	user, _ := NewUser(id, username, password)

	if user.CheckPassword(password) != true {
		t.Errorf("Expected password to match")
	}
	if user.CheckPassword("wrongpassword") != false {
		t.Errorf("Expected password to not match")
	}
}
