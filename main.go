package main

import (
	"fmt"
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
	userRepository := repository.NewUserRepositoryImpl(database)
	menuRepository := repository.NewMenuRepositoryImpl(database)

	// Service
	foodService := service.NewFoodServiceImpl(&foodRepository)
	userService := service.NewUserServiceImpl(&userRepository)
	menuService := service.NewMenuServiceImpl(&menuRepository)

	// Controller
	homeController := controller.NewHomeController()
	foodController := controller.NewFoodController(&foodService, config)
	userController := controller.NewUserController(&userService, config)
	menuController := controller.NewMenuController(&menuService, config)

	// Setup Fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	// Routing
	foodController.Route(app)
	homeController.Route(app)
	userController.Route(app)
	menuController.Route(app)

	log.Fatal(
		app.Listen(fmt.Sprintf(":%s", config.Get("SERVER_PORT"))),
	)
}
