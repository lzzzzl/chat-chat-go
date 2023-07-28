package repository

import (
	"database/sql"
	"errors"
	"github/lzzzzl/chat-chat-go/domain/model"
)

// UserRepository provides an interface to persist User domain struct object to the underlying database
type UserRepository interface {
	FindByID(id string) (*model.User, error)
	Save(user *model.User) error
	FindByUsername(username string) (*model.User, error)
}

// PostgresUserRepository Implementation
type PostgresUserRepository struct {
	DB *sql.DB
}

// NewPostgresUserRepository will create an object that represent the UserRepository interface
func NewPostgresUserRepository(DB *sql.DB) UserRepository {
	return &PostgresUserRepository{DB}
}

func (r *PostgresUserRepository) FindByID(id string) (*model.User, error) {
	row := r.DB.QueryRow("SELECT id, username, password_hash FROM user WHERE id = $1", id)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) Save(user *model.User) error {
	_, err := r.DB.Exec("INSERT INTO users (id, username, password_hash) VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE SET username = $2, password_hash = $3",
		user.ID, user.Username, user.PasswordHash)
	return err
}

func (r *PostgresUserRepository) FindByUsername(username string) (*model.User, error) {
	row := r.DB.QueryRow("SELECT id, username, password_hash FROM user WHERE username = $1", username)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
	}

	return user, nil
}
