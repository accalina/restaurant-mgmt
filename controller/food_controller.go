package controller

import (
	"github.com/accalina/restaurant-mgmt/configuration"
	"github.com/accalina/restaurant-mgmt/exception"
	"github.com/accalina/restaurant-mgmt/middleware"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/service"
	"github.com/gofiber/fiber/v2"
)

type FoodController struct {
	service.FoodServie
	configuration.Config
}

func NewFoodController(foodService *service.FoodServie, config configuration.Config) *FoodController {
	return &FoodController{FoodServie: *foodService, Config: config}
}

func (controller FoodController) Route(app *fiber.App) {
	food := app.Group("/food", middleware.LoggerMiddleware)
	food.Post("/", controller.Create)
}

func (controller FoodController) Create(c *fiber.Ctx) error {
	var request model.FoodCreteOrUpdateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.FoodServie.Create(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}
