package controller

import (
	"github.com/accalina/restaurant-mgmt/common"
	"github.com/accalina/restaurant-mgmt/middleware"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (homeController HomeController) Route(app *fiber.App) {
	home := app.Group("", middleware.LoggerMiddleware)
	home.Get("/", middleware.LoggerMiddleware, homeController.FindAll)
}

func (homeController HomeController) FindAll(c *fiber.Ctx) error {
	dt := common.DataMessageValue{Message: "Welcome to Restaurant Management API"}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    dt.GetDataMessage(),
	})
}
