package main

import (
	"log"

	"github.com/accalina/restaurant-mgmt/src"
	"github.com/accalina/restaurant-mgmt/src/models"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Connect to DB
	if err := models.ConnDB(); err != nil {
		log.Fatal("cannot connect to DB")
	}

	// Migrate the DB
	if err := models.DB.AutoMigrate(&models.FoodItem{}); err != nil {
		log.Fatal("cannot migrate DB")
	}

	app := fiber.New()

	src.Routes(app)

	log.Fatal(
		app.Listen(":8001"),
	)
}
