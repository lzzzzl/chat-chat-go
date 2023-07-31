package model

import "errors"

// Room represents a chat room
type Room struct {
	ID   int
	Name string
}

func NewRoom(name string) (*Room, error) {
	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	return &Room{
		Name: name,
	}, nil
}
