package db

import "testing"

func TestNewPostgresDB(t *testing.T) {
	db, err := NewPostgresDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
