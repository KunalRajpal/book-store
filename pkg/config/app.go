package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect() error {
    // Use DATABASE_URL environment variable provided by Heroku
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        dsn = "fallback:connection/string"  // Fallback for local development
    }

    d, err := gorm.Open("postgres", dsn)
    if err != nil {
        return err
    }
    db = d
    return nil
}

func GetDB() *gorm.DB {
    return db
}
