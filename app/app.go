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
	"github.com/accalina/restaurant-mgmt/platform/database"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	engine *fiber.App
}

func New() *App {
	app := &App{
		engine: fiber.New(configuration.NewFiberConfiguration()),
	}

	return app
}

func (a *App) Serve() {
	middleware.FiberMiddleware(a.engine)

	userController := controller.NewUserController(service.GetSharedService())
	menuController := controller.NewMenuController(service.GetSharedService())
	foodController := controller.NewFoodController(service.GetSharedService())
	tableController := controller.NewTableController(service.GetSharedService())
	orderController := controller.NewOrderController(service.GetSharedService())
	orderItemController := controller.NewOrderItemController(service.GetSharedService())
	invoiceController := controller.NewInvoiceController(service.GetSharedService())
	userController.Route(a.engine)
	menuController.Route(a.engine)
	foodController.Route(a.engine)
	tableController.Route(a.engine)
	orderController.Route(a.engine)
	orderItemController.Route(a.engine)
	invoiceController.Route(a.engine)

	log.Fatal(
		a.engine.Listen(fmt.Sprintf(":%s", env.BaseEnv().ServerPort)),
	)
}

func LoadAppConfig() {
	env.Load()
	DB := database.NewDatabase()
	configuration.NewRedis()
	repository.SetSharedRepoSQL(DB)
	service.SetSharedService()
}
