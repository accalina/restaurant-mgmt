package main

import (
	"log"

	"github.com/accalina/restaurant-mgmt/configuration"
	"github.com/accalina/restaurant-mgmt/controller"
	repository "github.com/accalina/restaurant-mgmt/repository/impl"
	service "github.com/accalina/restaurant-mgmt/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// Setup config
	isRunMigration := true
	config := configuration.New()
	database := configuration.NewDatabase(config, isRunMigration)

	//  Repository
	foodRepository := repository.NewFoodRepositoryImpl(database)

	// Service
	foodService := service.NewFoodServiceImpl(&foodRepository)

	// Controller
	foodController := controller.NewFoodController(&foodService, config)
	homeController := controller.NewHomeController()

	// Setup Fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	// Routing
	foodController.Route(app)
	homeController.Route(app)

	log.Fatal(
		app.Listen(config.Get("SERVER_PORT")),
	)
}
