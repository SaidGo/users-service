package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init открывает БД SQLite. Путь берётся из USERS_DB_DSN, по умолчанию users.db в рабочей директории.
func Init() {
	dsn := os.Getenv("USERS_DB_DSN")
	if dsn == "" {
		dsn = "users.db"
	}
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("open sqlite: %v", err)
	}
	DB = db
}
