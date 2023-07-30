package repository

import (
	"database/sql"
	"errors"
	"github/lzzzzl/chat-chat-go/domain/model"
)

// MessageRepository provides an interface to persist Message domain struct object to the underlying database
type MessageRepository interface {
	FindByID(id int) (*model.Message, error)
	Save(message *model.Message) error
	DeleteByID(id int) error
	FindMessagesByRoomID(roomID int) ([]*model.Message, error)
}

// PostgresMessageRepository Implementation
type PostgresMessageRepository struct {
	DB *sql.DB
}

// NewPostgresMessageRepository will create an object that represent the MessageRepository interface
func NewPostgresMessageRepository(DB *sql.DB) MessageRepository {
	return &PostgresMessageRepository{DB}
}

func (r *PostgresMessageRepository) FindByID(id int) (*model.Message, error) {
	row := r.DB.QueryRow("SELECT id, user_id, room_id, content, created_at FROM messages WHERE id = $1", id)

	message := &model.Message{}
	err := row.Scan(&message.ID, &message.UserID, &message.RoomID, &message.Content, &message.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("message not found")
		}
		return nil, err
	}

	return message, nil
}

func (r *PostgresMessageRepository) Save(message *model.Message) error {
	return r.DB.QueryRow("INSERT INTO messages (user_id, room_id, content, created_at) VALUES ($1, $2, $3, $4)  RETURNING id",
		message.UserID, message.RoomID, message.Content, message.CreatedAt).Scan(&message.ID)
}

func (r *PostgresMessageRepository) DeleteByID(id int) error {
	_, err := r.DB.Exec("DELETE FROM messages WHERE id = $1", id)
	return err
}

func (r *PostgresMessageRepository) FindMessagesByRoomID(roomID int) ([]*model.Message, error) {
	rows, err := r.DB.Query("SELECT id, user_id, room_id, content, created_at FROM messages WHERE room_id = $1", roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*model.Message
	for rows.Next() {
		message := &model.Message{}
		if err := rows.Scan(&message.ID, &message.UserID, &message.RoomID, &message.Content, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
