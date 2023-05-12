package controller

import (
	"github.com/accalina/restaurant-mgmt/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", middleware.LoggerMiddleware, Home)

	// Define a middleware function to log all request
	food := app.Group("/food", middleware.LoggerMiddleware)
	food.Get("/", ListFood)
	food.Get("/:id", RetrieveFood)
	food.Post("/", InsertFood)
	food.Put("/:id", UpdateFood)
	food.Patch("/:id", PatchFood)

}
