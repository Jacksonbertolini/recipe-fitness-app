package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Open initialises the connection pool. Call once from main().
func Open() {
	dsn := buildDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("database: failed to open: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("database: failed to connect: %v", err)
	}

	// Sensible pool settings for a small Lightsail instance
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	DB = db
	log.Println("database: connected to MySQL")
}

func buildDSN() string {
	host := getenv("DB_HOST", "localhost")
	port := getenv("DB_PORT", "3306")
	user := getenv("DB_USER", "fitmeals_user")
	pass := getenv("DB_PASSWORD", "fitmeals_password")
	name := getenv("DB_NAME", "fitmeals")

	// parseTime=true makes MySQL TIMESTAMP columns scan into time.Time
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
