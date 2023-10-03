package controller

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(app *fiber.App) {
	route := app.Group("/swagger")
	route.Get("*", swagger.New(swagger.Config{ // custom
		URL:         "/docs/swagger.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	}))
}
