package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the chat application.
type User struct {
	ID           int
	Username     string
	PasswordHash string
}

// NewUser creates a new User
// This function will also validate the input and hash the password.
func NewUser(username string, password string) (*User, error) {
	if username == "" {
		return nil, errors.New("username must not be empty")
	}
	if password == "" {
		return nil, errors.New("password must not be empty")
	}
	user := &User{
		Username: username,
	}
	err := user.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// hashPassword hashed a password using bcrypt.
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(bytes)
	return nil
}

// CheckPassword checks if the provided password matches the hashed password.
func (u *User) CheckPassword(plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(plainPassword))
	return err
}
