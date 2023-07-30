package repository

import (
	"github/lzzzzl/chat-chat-go/domain/model"
	"github/lzzzzl/chat-chat-go/infrastructure/db"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMessageRepository(t *testing.T) {
	db, err := db.NewPostgresDB()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPostgresMessageRepository(db)

	message, err := model.NewMessage(2, 1, "Test Content", time.Now())
	assert.NoError(t, err)

	t.Run("Save and FindByID", func(t *testing.T) {
		// Save the message
		err = repo.Save(message)
		assert.NoError(t, err)

		// Retrieve the message BY ID
		got, err := repo.FindByID(message.ID)
		assert.NoError(t, err)
		assert.Equal(t, message.ID, got.ID)
		assert.Equal(t, message.UserID, got.UserID)
		assert.Equal(t, message.RoomID, got.RoomID)
		assert.Equal(t, message.Content, got.Content)
	})

	t.Run("FindMessagesByRoomID", func(t *testing.T) {
		messages, err := repo.FindMessagesByRoomID(1)
		assert.NoError(t, err)

		assert.Greater(t, len(messages), 0)
	})

	t.Run("Delete", func(t *testing.T) {
		err := repo.DeleteByID(message.ID)
		assert.NoError(t, err)
	})
}
