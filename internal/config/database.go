package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type DBStatus struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	MaxConns  int    `json:"max_connections"`
	OpenConns int    `json:"open_connections"`
	InUse     int    `json:"in_use_connections"`
	Idle      int    `json:"idle_connections"`
}

func NewDBConnection() (*sql.DB, error) {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	fmt.Println(dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}

func CheckDBConnection(db *sql.DB) (*DBStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.PingContext(ctx)
	status := &DBStatus{
		MaxConns:  db.Stats().MaxOpenConnections,
		OpenConns: db.Stats().OpenConnections,
		InUse:     db.Stats().InUse,
		Idle:      db.Stats().Idle,
	}

	if err != nil {
		status.Status = "error"
		status.Message = fmt.Sprintf("Database connection failed: %v", err)
		return status, err
	}

	status.Status = "healthy"
	status.Message = "Database connection is healthy"
	return status, nil
} 