package model

import "errors"

// Room represents a chat room
type Room struct {
	ID   string
	Name string
}

func NewRoom(id string, name string) (*Room, error) {
	if id == "" {
		return nil, errors.New("id must not be empty")
	}
	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	return &Room{
		ID:   id,
		Name: name,
	}, nil
}
