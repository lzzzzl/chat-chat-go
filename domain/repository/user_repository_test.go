package repository

import (
	"github/lzzzzl/chat-chat-go/domain/model"
	"github/lzzzzl/chat-chat-go/infrastructure/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	db, err := db.NewPostgresDB()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPostgresUserRepository(db)

	user, err := model.NewUser("testuser1", "123")
	assert.NoError(t, err)

	t.Run("Save and FindByID", func(t *testing.T) {
		// Save the user
		err = repo.Save(user)
		assert.NoError(t, err)

		// Retrieve the user by ID
		got, err := repo.FindByID(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user, got)
	})

	t.Run("FindByUsername", func(t *testing.T) {
		got, err := repo.FindByUsername("testuser1")
		assert.NoError(t, err)
		assert.Equal(t, user, got)
	})

	t.Run("Delete", func(t *testing.T) {
		err := repo.DeleteByID(user.ID)
		assert.NoError(t, err)
	})

}
