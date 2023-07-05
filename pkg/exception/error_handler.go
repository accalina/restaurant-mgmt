package exception

import (
	"encoding/json"

	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	_, validationerror := err.(ValidationError)
	if validationerror {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		return c.Status(400).JSON(model.GeneralResponse{
			Success: false,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		return c.Status(404).JSON(model.GeneralResponse{
			Success: false,
			Message: "Not Found",
			Data:    messages,
		})
	}

	return c.Status(500).JSON(model.GeneralResponse{
		Success: false,
		Message: "General Error",
		Data:    err.Error(),
	})
}
