package controller

import (
	"fmt"

	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/pkg/common"
	"github.com/accalina/restaurant-mgmt/pkg/exception"
	"github.com/accalina/restaurant-mgmt/pkg/middleware"
	"github.com/accalina/restaurant-mgmt/pkg/shared/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct{ service service.Service }

func NewUserController(service service.Service) *UserController {
	return &UserController{service: service}
}

func (userController UserController) Route(app *fiber.App) {
	user := app.Group("/user", middleware.LoggerMiddleware)
	user.Post("/register", userController.Register)
	user.Post("/login", userController.Login)
	user.Post("/logout", middleware.RegisteredLogger, userController.Logout)
	user.Get("/", middleware.AdminLogger, userController.FindAll)
	user.Get("/:id", middleware.AdminLogger, userController.FindById)
	user.Delete("/:id", middleware.AdminLogger, userController.Delete)
}

func (userController UserController) FindAll(c *fiber.Ctx) error {
	response := userController.service.User().FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (userController UserController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	response, err := userController.service.User().FindById(c.Context(), id)
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

// @Description	Create a new User.
// @Summary		create a new User
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		User	body		model.UserCreateModel	true	"User attribute"
// @Success		200				{object}	model.ResponseLogin"
// @Security	ApiKeyAuth
// @Router		/user/register [post]
func (userController UserController) Register(c *fiber.Ctx) error {
	var request model.UserCreateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)
	err = common.Validate(&request)
	exception.PanicLogging(err)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request",
			Data:    err,
			Errors:  model.Extract(err),
		})
	}

	response := userController.service.User().Register(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Success",
		Data:    response,
		Success: true,
	})
}

// GetNewAccessToken method for create a new access token.
//
//	@Description	Create a new access token.
//	@Summary		create a new access token
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			loginRequest	body		model.UserCreateModel	true	"Login atributes"
//	@Success		200				{object}	model.ResponseLogin"
//	@Router			/user/login [post]
func (userController UserController) Login(c *fiber.Ctx) error {
	var request model.UserCreateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)
	err = common.Validate(&request)
	exception.PanicLogging(err)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request",
			Data:    err,
			Errors:  model.Extract(err),
		})
	}

	response, err := userController.service.User().Login(c.Context(), request.Username, request.Password)
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

func (userController UserController) Logout(c *fiber.Ctx) error {
	username := fmt.Sprintf("%s", c.Locals("username"))
	userController.service.User().Logout(c.Context(), username)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

func (userController UserController) Promote(c *fiber.Ctx) error {
	id := c.Params("id")

	deleteErr := userController.service.User().Promote(c.Context(), id)
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

	deleteErr := userController.service.User().Demote(c.Context(), id)
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
	err = common.Validate(&request)
	exception.PanicLogging(err)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request",
			Data:    err,
			Errors:  model.Extract(err),
		})
	}

	response := userController.service.User().Update(c.Context(), request, id)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (userController UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	deleteErr := userController.service.User().Delete(c.Context(), id)
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
