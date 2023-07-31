package repository

import (
	"github/lzzzzl/chat-chat-go/domain/model"
	"github/lzzzzl/chat-chat-go/infrastructure/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoomRepository(t *testing.T) {
	db, err := db.NewPostgresDB()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPostgresRoomRepository(db)

	room, err := model.NewRoom("chat room")
	assert.NoError(t, err)

	t.Run("Save and FindByID", func(t *testing.T) {
		// Save the room
		err = repo.Save(room)
		assert.NoError(t, err)

		// Retrieve the user by ID
		got, err := repo.FindByID(room.ID)
		assert.NoError(t, err)
		assert.Equal(t, room, got)
	})

	t.Run("FindAll", func(t *testing.T) {
		rooms, err := repo.FindAll()
		assert.NoError(t, err)
		assert.Greater(t, len(rooms), 0)
	})
}
