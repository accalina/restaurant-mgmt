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

type UserController struct {
	service.UserService
	configuration.Config
}

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

func (userController UserController) Route(app *fiber.App) {
	user := app.Group("/user", middleware.LoggerMiddleware)
	user.Post("/register", userController.Register)
	user.Post("/login", userController.Login)
	user.Get("/", userController.FindAll)
	user.Get("/:id", userController.FindById)
	user.Delete("/:id", userController.Delete)
}

func (userController UserController) FindAll(c *fiber.Ctx) error {
	response := userController.UserService.FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (userController UserController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	response, err := userController.UserService.FindById(c.Context(), id)
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

func (userController UserController) Register(c *fiber.Ctx) error {
	var request model.UserCreateModel
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

	response := userController.UserService.Register(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Success",
		Data:    response,
	})
}

func (userController UserController) Login(c *fiber.Ctx) error {
	var request model.UserCreateModel
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

	response, err := userController.UserService.Login(c.Context(), request.Username, request.Password)
	if err != nil {
		return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid username / password",
			Data:    err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Success",
		Data:    response,
	})
}

func (userController UserController) Promote(c *fiber.Ctx) error {
	id := c.Params("id")

	deleteErr := userController.UserService.Promote(c.Context(), id)
	if deleteErr {
		emptyArr := common.DataArrayValue{ArrMessage: []string{}}
		return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
			Code:    200,
			Message: "Error",
			Data:    emptyArr.ArrMessage,
		})
	}
	message := common.DataMessageValue{Message: "User ID: " + id + " has been promoted to admin!"}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    message.Message,
	})
}

func (userController UserController) Demote(c *fiber.Ctx) error {
	id := c.Params("id")

	deleteErr := userController.UserService.Demote(c.Context(), id)
	if deleteErr {
		emptyArr := common.DataArrayValue{ArrMessage: []string{}}
		return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
			Code:    200,
			Message: "Error",
			Data:    emptyArr.ArrMessage,
		})
	}
	message := common.DataMessageValue{Message: "User ID: " + id + " has been promoted to admin!"}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    message.Message,
	})
}

func (userController UserController) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var request model.UserUpdateModel
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

	response := userController.UserService.Update(c.Context(), request, id)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (userController UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	deleteErr := userController.UserService.Delete(c.Context(), id)
	if deleteErr {
		emptyArr := common.DataArrayValue{ArrMessage: []string{}}
		return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
			Code:    200,
			Message: "Error",
			Data:    emptyArr.ArrMessage,
		})
	}
	message := common.DataMessageValue{Message: "User ID: " + id + " has been deleted!"}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    message.Message,
	})
}
