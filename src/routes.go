package src

import (
	"github.com/accalina/restaurant-mgmt/src/middlewares"
	"github.com/accalina/restaurant-mgmt/src/pages"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", pages.Home)

	// Define a middleware function to log all request
	food := app.Group("/food", middlewares.LoggerMiddleware)
	food.Get("/", pages.ListFood)
	food.Get("/:id", pages.RetrieveFood)
	food.Post("/", pages.InsertFood)
	food.Put("/:id", pages.UpdateFood)
	food.Patch("/:id", pages.PatchFood)

}
