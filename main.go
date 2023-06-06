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
	tableRepository := repository.NewTableRepositoryImpl(database)
	orderRepository := repository.NewOrderRepositoryImpl(database)
	orderItemRepository := repository.NewOrderItemRepositoryImpl(database)
	invoiceRepository := repository.NewInvoiceRepositoryImpl(database)

	// Service
	foodService := service.NewFoodServiceImpl(&foodRepository)
	userService := service.NewUserServiceImpl(&userRepository)
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
