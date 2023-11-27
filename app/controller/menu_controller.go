package controller

import (
	"fmt"

	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/pkg/common"
	"github.com/accalina/restaurant-mgmt/pkg/middleware"
	"github.com/accalina/restaurant-mgmt/pkg/shared/service"
	"github.com/gofiber/fiber/v2"
)

type MenuController struct{ service service.Service }

func NewMenuController(service service.Service) *MenuController {
	return &MenuController{service: service}
}

func (c MenuController) Route(app *fiber.App) {
	menu := app.Group("/menu", middleware.LoggerMiddleware)
	menu.Get("/", c.getAllMenu)
	menu.Get("/:id", c.getDetailMenuByID)
	menu.Post("/", middleware.AdminLogger, c.createMenu)
	menu.Put("/:id", middleware.AdminLogger, c.updateMenu)
	menu.Delete("/:id", middleware.AdminLogger, c.deleteMenu)
}

// @Summary			List all menu
// @Description		List all menu.
// @Tags			Menu
// @Accept			json
// @Produce			json
// @Param        	search	query     string  false  "name search"
// @Param        	limit	query     string  false  "limit search"
// @Param        	page    query     string  false  "page search"
// @Success			200		{array}		entity.Menu
// @Failure			400		{object}	model.GeneralResponse
// @Failure			500		{object}	model.GeneralResponse
// @Router			/menu [get]
func (c MenuController) getAllMenu(ctx *fiber.Ctx) error {
	queryParams := model.NewMenuFilter("Foods")
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid parameter",
			Data:    err,
		})
	}

	response, meta, err := c.service.Menu().GetAllMenu(ctx.Context(), queryParams)
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

// @Summary			Detail menu
// @Description		Detail menu.
// @Tags			Menu
// @Accept			json
// @Produce			json
// @Success			200		{object}	entity.Menu
// @Failure			400		{object}	model.GeneralResponse
// @Failure			500		{object}	model.GeneralResponse
// @Router			/menu/{id} [get]
func (c *MenuController) getDetailMenuByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	menu, err := c.service.Menu().GetDetailMenu(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Get detail menu failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    menu,
	})
}

// @Description	Create a new menu.
// @Summary		create a new menu
// @Tags		Menu
// @Accept		json
// @Produce		json
// @Param		Menu	body		model.MenuCreateOrUpdateSwaggerModel	true	"Menu attribute"
// @Success		200		{object}	entity.Menu
// @Failure		400		{object}	model.GeneralResponse
// @Failure		500		{object}	model.GeneralResponse
// @Security	ApiKeyAuth
// @Router		/menu [post]
func (c *MenuController) createMenu(ctx *fiber.Ctx) error {
	var request model.MenuCreateOrUpdateModel
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

	result, err := c.service.Menu().CreateMenu(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Create Menu failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

// @Description	Update menu.
// @Summary		Update menu
// @Tags		Menu
// @Accept		json
// @Produce		json
// @Param		Menu	body		model.MenuCreateOrUpdateSwaggerModel	true	"Menu attribute"
// @Success		200		{object}	entity.Menu
// @Failure		400		{object}	model.GeneralResponse
// @Failure		500		{object}	model.GeneralResponse
// @Security	ApiKeyAuth
// @Router		/menu/{id} [put]
func (c *MenuController) updateMenu(ctx *fiber.Ctx) error {
	var request model.MenuCreateOrUpdateModel
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
	result, err := c.service.Menu().UpdateMenu(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Update Menu failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

// @Description	Delete menu.
// @Summary		Delete menu
// @Tags		Menu
// @Accept		json
// @Produce		json
// @Success		200		{object}	entity.Menu
// @Failure		400		{object}	model.GeneralResponse
// @Failure		500		{object}	model.GeneralResponse
// @Security	ApiKeyAuth
// @Router		/menu/{id} [delete]
func (c *MenuController) deleteMenu(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.service.Menu().DeleteMenu(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Delete Menu failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: fmt.Sprintf("Menu ID: %s has been deleted", id),
	})
}
