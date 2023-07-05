package controller

import (
	"fmt"

	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/pkg/common"
	"github.com/accalina/restaurant-mgmt/pkg/middleware"
	"github.com/accalina/restaurant-mgmt/pkg/shared/service"
	"github.com/gofiber/fiber/v2"
)

type TableController struct{ service service.Service }

func NewTableController(service service.Service) *TableController {
	return &TableController{service: service}
}

func (c TableController) Route(app *fiber.App) {
	table := app.Group("/table", middleware.LoggerMiddleware)
	table.Get("/", c.getAllTable)
	table.Get("/:id", c.getDetailTableByID)
	table.Post("/", c.createTable)
	table.Put("/:id", c.updateTable)
	table.Delete("/:id", c.deleteTable)
}

func (c TableController) getAllTable(ctx *fiber.Ctx) error {
	queryParams := model.NewTableFilter()
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid parameter",
			Data:    err,
		})
	}

	response, meta, err := c.service.Table().GetAllTable(ctx.Context(), queryParams)
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

func (c *TableController) getDetailTableByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	table, err := c.service.Table().GetDetailTable(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Get detail table failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    table,
	})
}

func (c *TableController) createTable(ctx *fiber.Ctx) error {
	var request model.TableCreateOrUpdateModel
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

	result, err := c.service.Table().CreateTable(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Create Table failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

func (c *TableController) updateTable(ctx *fiber.Ctx) error {
	var request model.TableCreateOrUpdateModel
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
	result, err := c.service.Table().UpdateTable(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Update Table failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

func (c *TableController) deleteTable(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.service.Table().DeleteTable(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Delete Table failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: fmt.Sprintf("Table ID: %s has been deleted", id),
	})
}
