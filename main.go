package main

import (
	"log"

	"github.com/accalina/restaurant-mgmt/configuration"
	"github.com/accalina/restaurant-mgmt/controller"
	repository "github.com/accalina/restaurant-mgmt/repository/impl"
	service "github.com/accalina/restaurant-mgmt/service/impl"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Setup config
	config := configuration.New()
	database := configuration.NewDatabase(config)

	//  Repository
	foodRepository := repository.NewFoodRepositoryImpl(database)

	// Service
	foodService := service.NewFoodServiceImpl(&foodRepository)

	// Controller
	foodController := controller.NewFoodController(&foodService, config)
	homeController := controller.NewHomeController()

	// Migrate the DB
	// if err := configuration.DB.AutoMigrate(&model.FoodItem{}); err != nil {
	// 	log.Fatal("cannot migrate DB")
	// }

	// Setup Fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	// app.Use(recover.New())
	// app.Use(cors.New())

	// Routing
	foodController.Route(app)
	homeController.Route(app)

	log.Fatal(
		app.Listen(config.Get("SERVER_PORT")),
	)
}
