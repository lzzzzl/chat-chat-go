package model

import (
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {
	username := "test"
	password := "password"

	user, err := NewUser(username, password)
	fmt.Printf("password, %s", user.PasswordHash)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if user.Username != username {
		t.Errorf("Expected Username to be %s, got %s", username, user.Username)
	}
	if user.CheckPassword(password) != true {
		t.Errorf("Expected password to match")
	}
}

func TestCheckPassword(t *testing.T) {
	username := "test"
	password := "password"
	user, _ := NewUser(username, password)

	if user.CheckPassword(password) != true {
		t.Errorf("Expected password to match")
	}
	if user.CheckPassword("wrongpassword") != false {
		t.Errorf("Expected password to not match")
	}
}
