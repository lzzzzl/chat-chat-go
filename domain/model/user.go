package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the chat application.
type User struct {
	ID           string
	Username     string
	PasswordHash string
}

// NewUser creates a new User
// This function will also validate the input and hash the password.
func NewUser(id string, username string, password string) (*User, error) {
	if id == "" {
		return nil, errors.New("id must not be empty")
	}
	if username == "" {
		return nil, errors.New("username must not be empty")
	}
	if password == "" {
		return nil, errors.New("password must not be empty")
	}
	// Hash the password
	passwordHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
	}, nil
}

// hashPassword hashed a password using bcrypt.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword checks if the provided password matches the hashed password.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
