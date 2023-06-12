package main

import (
	"fmt"
	"log"

	"github.com/accalina/restaurant-mgmt/configuration"
	"github.com/accalina/restaurant-mgmt/controller"
	_ "github.com/accalina/restaurant-mgmt/docs"
	repository "github.com/accalina/restaurant-mgmt/repository/impl"
	service "github.com/accalina/restaurant-mgmt/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title RESTAURANT-MGMT API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {

	// Setup config
	isRunMigration := true
	config := configuration.New()
	database := configuration.NewDatabase(config, isRunMigration)
	redis := configuration.NewRedis(config)

	//  Repository
	foodRepository := repository.NewFoodRepositoryImpl(database)
	userRepository := repository.NewUserRepositoryImpl(database)
	menuRepository := repository.NewMenuRepositoryImpl(database)
	tableRepository := repository.NewTableRepositoryImpl(database)
	orderRepository := repository.NewOrderRepositoryImpl(database)
	orderItemRepository := repository.NewOrderItemRepositoryImpl(database)
	invoiceRepository := repository.NewInvoiceRepositoryImpl(database)

	// Service
	foodService := service.NewFoodServiceImpl(&foodRepository)
	userService := service.NewUserServiceImpl(&userRepository, redis)
	menuService := service.NewMenuServiceImpl(&menuRepository)
	tableService := service.NewTableServiceImpl(&tableRepository)
	orderService := service.NewOrderServiceImpl(&orderRepository)
	orderItemService := service.NewOrderItemServiceImpl(&orderItemRepository)
	InvoiceService := service.NewInvoiceServiceImpl(&invoiceRepository)

	// Controller
	homeController := controller.NewHomeController()
	foodController := controller.NewFoodController(&foodService, config)
	userController := controller.NewUserController(&userService, config)
	menuController := controller.NewMenuController(&menuService, config)
	tableController := controller.NewTableController(&tableService, config)
	orderController := controller.NewOrderController(&orderService, config)
	orderItemController := controller.NewOrderItemController(&orderItemService, config)
	invoiceController := controller.NewInvoiceController(&InvoiceService, config)

	// Setup Fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	// Routing
	controller.SwaggerRoute(app) // Register a route for API Docs (Swagger).
	foodController.Route(app)
	homeController.Route(app)
	userController.Route(app)
	menuController.Route(app)
	tableController.Route(app)
	orderController.Route(app)
	orderItemController.Route(app)
	invoiceController.Route(app)

	log.Fatal(
		app.Listen(fmt.Sprintf(":%s", config.Get("SERVER_PORT"))),
	)
}
