package db

import (
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"

	// 这里导入lib/pq包，即使我们在这个文件中并没有直接使用它
	_ "github.com/lib/pq"
)

type Config struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func NewPostgresDB() (*sql.DB, error) {
	configFile := "../../config.toml"
	config := Config{}
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.Postgres.User, config.Postgres.Password, config.Postgres.Host, config.Postgres.Dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
