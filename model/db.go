package model

// import (
// 	"fmt"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnDB() error {
// 	// Load the .env vars
// 	if err := godotenv.Load(); err != nil {
// 		return err
// 	}

// 	// Construct the DSN string
// 	dsn := fmt.Sprintf(
// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
// 		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"),
// 	)

// 	// Open new connection to PostgreSQL
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}

// 	DB = db
// 	return nil
// }
