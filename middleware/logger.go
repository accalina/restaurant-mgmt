package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	log.Printf("Request: %s %s\n", c.Method(), c.Path())
	return c.Next()
}
