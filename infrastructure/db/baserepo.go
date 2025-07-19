package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Config struct {
	User string `json:"db_user"`
	Pass string `json:"db_pass"`
	Host string `json:"db_host"`
	Port int    `json:"db_port"`
	Name string `json:"db_name"`
	SSL  string `json:"db_ssl"`
}

type DB struct {
	Conn *sql.DB
}

func loadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}

func New(configPath string) (*DB, error) {
	cfg, err := loadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.Name,
		cfg.SSL)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	return &DB{Conn: db}, nil
}

func InitializeDB(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS tasks_list (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            status BOOLEAN DEFAULT FALSE
        )
    `)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return err
	}
	return nil
}
