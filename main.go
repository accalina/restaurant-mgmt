package main

import (
	"log"

	"github.com/accalina/restaurant-mgmt/configuration"
	"github.com/accalina/restaurant-mgmt/controller"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// Setup config
	config := configuration.New()
	configuration.NewDatabase(config)

	// Migrate the DB
	if err := configuration.DB.AutoMigrate(&model.FoodItem{}); err != nil {
		log.Fatal("cannot migrate DB")
	}

	// Setup Fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	controller.Routes(app)

	log.Fatal(
		app.Listen(":8001"),
	)
}
