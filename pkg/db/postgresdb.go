package db

import (
	"OnlyGo/pkg/quote"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost port=5432 user=postgres password=yourpassword dbname=postgres sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("FATAL: Ошибка подключения к БД: %v", err)
	}
	if err := db.AutoMigrate(&quote.Quote{}); err != nil {
		log.Fatalf("FATAL: Ошибка миграции: %v", err)
	}

	return db, nil
}

