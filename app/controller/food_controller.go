package controller

import (
	"fmt"

	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/pkg/common"
	"github.com/accalina/restaurant-mgmt/pkg/middleware"
	"github.com/accalina/restaurant-mgmt/pkg/shared/service"
	"github.com/gofiber/fiber/v2"
)

type FoodController struct{ service service.Service }

func NewFoodController(service service.Service) *FoodController {
	return &FoodController{service: service}
}

func (foodController FoodController) Route(app *fiber.App) {
	food := app.Group("/food", middleware.LoggerMiddleware)
	food.Get("/", foodController.FindAll)
	food.Get("/:id", foodController.FindById)
	food.Post("/", middleware.AdminLogger, foodController.Create)
	food.Put("/:id", middleware.AdminLogger, foodController.Update)
	food.Delete("/:id", middleware.AdminLogger, foodController.Delete)
}

// @Summary			List all food
// @Description		List all food.
// @Tags			Food
// @Accept			json
// @Produce			json
// @Param        	search	query     string  false  "name search"
// @Param        	limit	query     string  false  "limit search"
// @Param        	page	query     string  false  "page search"
// @Success			200		{array}		entity.Food
// @Failure			400		{object}	model.GeneralResponse
// @Failure			500		{object}	model.GeneralResponse
// @Router			/food [get]
func (c FoodController) FindAll(ctx *fiber.Ctx) error {
	queryParams := model.NewFoodFilter("Menu")
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid parameter",
			Data:    err,
		})
	}

	response, meta, err := c.service.Food().GetAllFood(ctx.Context(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Fetch data failed: %s", err.Error()),
			Data:    response,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
		Meta:    meta,
	})
}

// @Summary			List detail food
// @Description		List detail food.
// @Tags			Food
// @Accept			json
// @Produce			json
// @Success			200		{object}		entity.Food
// @Router			/food/{id} [get]
func (c FoodController) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	food, err := c.service.Food().GetDetailFood(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Get detail food failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    food,
	})
}

// @Description	Create a new Food.
// @Summary		create a new Food
// @Tags		Food
// @Accept		json
// @Produce		json
// @Param		Food	body		model.FoodCreateOrUpdateSwaggerModel	true	"Food attribute"
// @Success		200		{object}	entity.Food
// @Failure		400		{object}	model.GeneralResponse
// @Failure		500		{object}	model.GeneralResponse
// @Security	ApiKeyAuth
// @Router		/food [post]
func (c FoodController) Create(ctx *fiber.Ctx) error {
	var request model.FoodCreateOrUpdateModel
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
	}
	if err = common.Validate(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request",
			Errors:  model.Extract(err),
		})
	}

	result, err := c.service.Food().CreateFood(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Create Food failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

// @Description	Update Food.
// @Summary		Update Food
// @Tags		Food
// @Accept		json
// @Produce		json
// @Param		Food	body		model.FoodCreateOrUpdateSwaggerModel	true	"Food attribute"
// @Success		200		{object}	entity.Food
// @Failure		400		{object}	model.GeneralResponse
// @Failure		500		{object}	model.GeneralResponse
// @Security	ApiKeyAuth
// @Router		/food/{id} [put]
func (c FoodController) Update(ctx *fiber.Ctx) error {
	var request model.FoodCreateOrUpdateModel
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
	}
	if err := common.Validate(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request",
			Errors:  model.Extract(err),
		})
	}

	request.ID = ctx.Params("id")
	result, err := c.service.Food().UpdateFood(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Update Food failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

// @Description	Delete Food.
// @Summary		Delete Food
// @Tags		Food
// @Accept		json
// @Produce		json
// @Success		200		{object}	model.GeneralResponse
// @Failure		400		{object}	model.GeneralResponse
// @Failure		500		{object}	model.GeneralResponse
// @Security	ApiKeyAuth
// @Router		/food/{id} [delete]
func (foodController FoodController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	deleteErr := foodController.service.Food().DeleteFood(c.Context(), id)
	if deleteErr != nil {
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
