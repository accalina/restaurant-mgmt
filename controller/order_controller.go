package controller

import (
	"fmt"

	"github.com/accalina/restaurant-mgmt/common"
	"github.com/accalina/restaurant-mgmt/configuration"
	"github.com/accalina/restaurant-mgmt/middleware"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/service"
	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	service.OrderService
	configuration.Config
}

func NewOrderController(s *service.OrderService, c configuration.Config) *OrderController {
	return &OrderController{OrderService: *s, Config: c}
}

func (c OrderController) Route(app *fiber.App) {
	order := app.Group("/order", middleware.RegisteredLogger)
	order.Get("/", c.getAllFood)
	order.Get("/:id", c.getDetailOrderByID)
	order.Post("/", c.createOrder)
	order.Put("/:id", c.updateOrder)
	order.Delete("/:id", c.deleteOrder)
}

func (c OrderController) getAllFood(ctx *fiber.Ctx) error {
	queryParams := model.NewOrderFilter()
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid parameter",
			Data:    err,
		})
	}

	response, meta, err := c.OrderService.GetAllOrder(ctx.Context(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Fetch data failed: %s", err.Error()),
			Data:    response,
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
		Meta:    meta,
	})
}

func (c *OrderController) getDetailOrderByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	order, err := c.OrderService.GetDetailOrder(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Get detail order failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    order,
	})
}

func (c *OrderController) createOrder(ctx *fiber.Ctx) error {
	var request model.OrderCreateOrUpdateModel
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

	result, err := c.OrderService.CreateOrder(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Create Order failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

func (c *OrderController) updateOrder(ctx *fiber.Ctx) error {
	var request model.OrderCreateOrUpdateModel
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
	result, err := c.OrderService.UpdateOrder(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Update Order failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

func (c *OrderController) deleteOrder(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.OrderService.DeleteOrder(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Delete Order failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: fmt.Sprintf("Order ID: %s has been deleted", id),
	})
}
