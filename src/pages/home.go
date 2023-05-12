package pages

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "welcome to restaurant management api",
		"success": true,
	})
}
