package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	if err := user.CheckPassword(password); err != nil {
		t.Errorf("Expected password to match")
	}
}

func TestCheckPassword(t *testing.T) {
	username := "test"
	password := "password"
	user, err := NewUser(username, password)
	assert.NoError(t, err)

	if err := user.CheckPassword(password); err != nil {
		t.Errorf("Expected password to match")
	}
	if err := user.CheckPassword("wrongpassword"); err == nil {
		t.Errorf("Expected password to not match")
	}
}
