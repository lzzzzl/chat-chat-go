package db

import (
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func TestMigrate(t *testing.T) {
	// 先回滚到最初的状态
	err := Rollback()
	if err != nil {
		t.Fatalf("Failed to rollback: %v", err)
	}

	// 执行迁移
	err = Migrate()
	if err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}
}

// Rollback 将数据库迁移回初始状态
func Rollback() error {
	db, err := NewPostgresDB()
	if err != nil {
		return err
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../scripts/",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
