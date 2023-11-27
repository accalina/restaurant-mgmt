package app

import (
	"fmt"
	"log"

	"github.com/accalina/restaurant-mgmt/app/controller"
	"github.com/accalina/restaurant-mgmt/pkg/configuration"
	"github.com/accalina/restaurant-mgmt/pkg/env"
	"github.com/accalina/restaurant-mgmt/pkg/middleware"
	"github.com/accalina/restaurant-mgmt/pkg/shared/repository"
	"github.com/accalina/restaurant-mgmt/pkg/shared/service"
	"github.com/accalina/restaurant-mgmt/platform/cache"
	"github.com/accalina/restaurant-mgmt/platform/database"
	"github.com/gofiber/fiber/v2"
)

type AppInstance struct {
	app *fiber.App
}

func New() *AppInstance {
	return &AppInstance{
		app: fiber.New(configuration.NewFiberConfiguration()),
	}
}

func (ai *AppInstance) Serve() {
	app := ai.app
	middleware.FiberMiddleware(app)

	// app.Static("/docs", "./docs")
	// app.Static("/swagger/docs.json", "./docs/swagger.json")

	userController := controller.NewUserController(service.GetSharedService())
	menuController := controller.NewMenuController(service.GetSharedService())
	foodController := controller.NewFoodController(service.GetSharedService())
	tableController := controller.NewTableController(service.GetSharedService())
	orderController := controller.NewOrderController(service.GetSharedService())
	orderItemController := controller.NewOrderItemController(service.GetSharedService())
	invoiceController := controller.NewInvoiceController(service.GetSharedService())

	userController.Route(app)
	menuController.Route(app)
	foodController.Route(app)
	tableController.Route(app)
	orderController.Route(app)
	orderItemController.Route(app)
	invoiceController.Route(app)
	controller.SwaggerRoute(app)

	app.Static("/docs", "./docs")
	log.Fatal(
		app.Listen(fmt.Sprintf(":%s", env.BaseEnv().ServerPort)),
	)
}

func LoadAppConfig() {
	env.Load()
	DB := database.NewDatabase()
	cache.NewRedis()
	repository.SetSharedRepoSQL(DB)
	service.SetSharedService()
}
