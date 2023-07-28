package db

import (
	"database/sql"
	"fmt"
)

func NewPostgresDB(source string) (*sql.DB, error) {
	db, err := sql.Open("postgres", source)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
