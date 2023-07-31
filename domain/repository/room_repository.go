package repository

import (
	"database/sql"
	"errors"
	"github/lzzzzl/chat-chat-go/domain/model"
)

// RoomRepository provides an interface to persist User domain struct object to the underlying database
type RoomRepository interface {
	Save(room *model.Room) error
	FindByID(id int) (*model.Room, error)
	FindAll() ([]*model.Room, error)
}

// PostgresRoomRepository Implementation
type PostgresRoomRepository struct {
	DB *sql.DB
}

// NewPostgresRoomRepository will create an object that represent the RoomRepository interface
func NewPostgresRoomRepository(DB *sql.DB) RoomRepository {
	return &PostgresRoomRepository{DB}
}

func (r *PostgresRoomRepository) Save(room *model.Room) error {
	return r.DB.QueryRow("INSERT INTO rooms (name) VALUES ($1) RETURNING id", room.Name).Scan(&room.ID)
}

func (r *PostgresRoomRepository) FindByID(id int) (*model.Room, error) {
	row := r.DB.QueryRow("SELECT id, name FROM rooms WHERE id = $1", id)

	room := &model.Room{}
	err := row.Scan(&room.ID, &room.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Room not found")
		}
		return nil, err
	}

	return room, nil
}

func (r *PostgresRoomRepository) FindAll() ([]*model.Room, error) {
	rows, err := r.DB.Query("SELECT id, name FROM rooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []*model.Room{}
	for rows.Next() {
		room := &model.Room{}
		err = rows.Scan(&room.ID, &room.Name)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
