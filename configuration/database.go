package configuration

import (
	"fmt"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDatabase(config Config, isRunMigration bool) *gorm.DB {
	// Construct the DSN string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.Get("DB_HOST"), config.Get("DB_PORT"), config.Get("DB_USER"), config.Get("DB_PASS"), config.Get("DB_NAME"),
	)

	// Open new connection to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exception.PanicLogging(err)

	// Run Migration
	// Related Model will be automatically migrated
	if isRunMigration {
		exception.PanicLogging(db.AutoMigrate(&entity.Food{}))
		exception.PanicLogging(db.AutoMigrate(&entity.OrderItem{}))
		exception.PanicLogging(db.AutoMigrate(&entity.Invoice{}))
		exception.PanicLogging(db.AutoMigrate(&entity.User{}))
	}

	return db
}
