package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrderModel struct {
	ID         string `gorm:"primaryKey"`
	Price      float64
	Tax        float64
	FinalPrice float64
	CreatedAt  time.Time
}

func Connect() *gorm.DB {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "host=db user=postgres password=postgres dbname=orders sslmode=disable"
	}

	// Retry connection to ensure database is ready
	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database, retrying in 3 seconds... (%s)", err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	// Automatically create the table if it doesn't exist
	if err := db.AutoMigrate(&OrderModel{}); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}

	return db
}
