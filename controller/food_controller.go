package controller

import (
	"github.com/accalina/restaurant-mgmt/common"
	"github.com/accalina/restaurant-mgmt/configuration"
	"github.com/accalina/restaurant-mgmt/exception"
	"github.com/accalina/restaurant-mgmt/middleware"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/service"
	"github.com/gofiber/fiber/v2"
)

type FoodController struct {
	service.FoodService
	configuration.Config
}

func NewFoodController(foodService *service.FoodService, config configuration.Config) *FoodController {
	return &FoodController{FoodService: *foodService, Config: config}
}

func (foodController FoodController) Route(app *fiber.App) {
	food := app.Group("/food", middleware.LoggerMiddleware)
	food.Get("/", foodController.FindAll)
	food.Get("/:id", foodController.FindById)
	food.Post("/", foodController.Create)
	food.Delete("/:id", foodController.Delete)
}

func (foodController FoodController) FindAll(c *fiber.Ctx) error {
	response := foodController.FoodService.FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (foodController FoodController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	response, err := foodController.FoodService.FindById(c.Context(), id)
	if err != nil {
		emptyArr := common.DataArrayValue{ArrMessage: []string{}}
		return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
			Code:    200,
			Message: "Error",
			Data:    emptyArr.ArrMessage,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (foodController FoodController) Create(c *fiber.Ctx) error {
	var request model.FoodCreteOrUpdateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)
	err = request.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request",
			Data:    err,
			Errors:  model.Extract(err),
		})
	}

	response := foodController.FoodService.Create(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Success",
		Data:    response,
	})
}

func (foodController FoodController) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var request model.FoodCreteOrUpdateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)
	err = request.Validate()
	exception.PanicLogging(err)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request",
			Data:    err,
			Errors:  model.Extract(err),
		})
	}

	response := foodController.FoodService.Update(c.Context(), request, id)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (foodController FoodController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	deleteErr := foodController.FoodService.Delete(c.Context(), id)
	if deleteErr {
		emptyArr := common.DataArrayValue{ArrMessage: []string{}}
		return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
			Code:    200,
			Message: "Error",
			Data:    emptyArr.ArrMessage,
		})
	}
	message := common.DataMessageValue{Message: "Food ID: " + id + " has been deleted!"}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    message.Message,
	})
}
