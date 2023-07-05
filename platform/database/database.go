package database

import (
	"fmt"
	"strings"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/pkg/env"
	"github.com/accalina/restaurant-mgmt/pkg/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	// Construct the DSN string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.BaseEnv().DBHost, env.BaseEnv().DBPort, env.BaseEnv().DBUser, env.BaseEnv().DBPass, env.BaseEnv().DBName,
	)

	// Open new connection to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exception.PanicLogging(err)

	// Run Migration
	// Related Model will be automatically migrated
	if strings.ToLower(env.BaseEnv().IsRunMigration) == "true" {
		exception.PanicLogging(db.AutoMigrate(&entity.Food{}))
		exception.PanicLogging(db.AutoMigrate(&entity.OrderItem{}))
		exception.PanicLogging(db.AutoMigrate(&entity.Invoice{}))
		exception.PanicLogging(db.AutoMigrate(&entity.User{}))
	}

	return db
}
