package controller

import (
	"fmt"

	"github.com/accalina/restaurant-mgmt/configuration"
	"github.com/accalina/restaurant-mgmt/middleware"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/service"
	"github.com/gofiber/fiber/v2"
)

type InvoiceController struct {
	service.InvoiceService
	configuration.Config
}

func NewInvoiceController(s *service.InvoiceService, c configuration.Config) *InvoiceController {
	return &InvoiceController{InvoiceService: *s, Config: c}
}

func (c InvoiceController) Route(app *fiber.App) {
	invoice := app.Group("/invoice", middleware.LoggerMiddleware)
	invoice.Get("/", c.getAllFood)
	invoice.Get("/:id", c.getDetailInvoiceByID)
	invoice.Post("/", c.createInvoice)
	invoice.Put("/:id", c.updateInvoice)
	invoice.Delete("/:id", c.deleteInvoice)
}

func (c InvoiceController) getAllFood(ctx *fiber.Ctx) error {
	queryParams := model.NewInvoiceFilter()
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid parameter",
			Data:    err,
		})
	}

	response, meta, err := c.InvoiceService.GetAllInvoice(ctx.Context(), queryParams)
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

func (c *InvoiceController) getDetailInvoiceByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	invoice, err := c.InvoiceService.GetDetailInvoice(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Get detail invoice failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    invoice,
	})
}

func (c *InvoiceController) createInvoice(ctx *fiber.Ctx) error {
	var request model.InvoiceCreateOrUpdateModel
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
	}
	if err = request.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request",
			Errors:  model.Extract(err),
		})
	}

	result, err := c.InvoiceService.CreateInvoice(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Create Invoice failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

func (c *InvoiceController) updateInvoice(ctx *fiber.Ctx) error {
	var request model.InvoiceCreateOrUpdateModel
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
	}
	if err := request.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request",
			Errors:  model.Extract(err),
		})
	}

	request.ID = ctx.Params("id")
	result, err := c.InvoiceService.UpdateInvoice(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Update Invoice failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

func (c *InvoiceController) deleteInvoice(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.InvoiceService.DeleteInvoice(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Delete Invoice failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: fmt.Sprintf("Invoice ID: %s has been deleted", id),
	})
}
