package middleware

import (
	"errors"
	"log"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/pkg/common"
	"github.com/accalina/restaurant-mgmt/pkg/configuration"
	"github.com/accalina/restaurant-mgmt/pkg/exception"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	log.Printf("IP: %s Request: %s %s\n", c.IP(), c.Method(), c.Path())
	return c.Next()
}

func RegisteredLogger(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	username, role, err := common.ParseJwt(auth)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
			Code:    401,
			Message: "Unauthorized",
			Success: false,
			Errors:  err.Error(),
		})
	}
	redisRole, err := configuration.RedisCache.Get(c.Context(), username).Result()
	if err == redis.Nil {
		redisRole = ""
	} else if err != nil {
		exception.PanicLogging(err)
	}

	if redisRole != role {
		return c.Status(403).JSON(model.GeneralResponse{
			Code:    403,
			Message: "Forbidden",
			Success: false,
			Errors:  errors.New("jwt is expired").Error(),
		})
	}
	c.Locals("username", username)
	c.Locals("role", role)
	log.Printf("IP: %s Request: %s %s\n", c.IP(), c.Method(), c.Path())
	return c.Next()
}

func AdminLogger(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	username, role, err := common.ParseJwt(auth)
	if err != nil {
		return c.Status(401).JSON(model.GeneralResponse{
			Code:    401,
			Message: "Unauthorized",
			Success: false,
			Errors:  err.Error(),
		})
	}

	redisRole, err := configuration.RedisCache.Get(c.Context(), username).Result()
	if err == redis.Nil {
		redisRole = ""
	} else if err != nil {
		exception.PanicLogging(err)
	}

	if redisRole != role {
		return c.Status(403).JSON(model.GeneralResponse{
			Code:    403,
			Message: "Forbidden",
			Success: false,
			Errors:  errors.New("jwt is expired").Error(),
		})
	}

	if role != entity.Role.Admin {
		return c.Status(403).JSON(model.GeneralResponse{
			Code:    403,
			Message: "Forbidden",
			Success: false,
			Errors:  errors.New("forbidden, only admin can enter").Error(),
		})
	}
	log.Printf("IP: %s Request: %s %s\n", c.IP(), c.Method(), c.Path())
	c.Locals("username", username)
	c.Locals("role", role)
	return c.Next()
}
