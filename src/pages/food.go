package pages

import (
	"strconv"

	"github.com/accalina/restaurant-mgmt/src/models"
	"github.com/gofiber/fiber/v2"
)

func ListFood(c *fiber.Ctx) error {
	var foods []models.FoodItem
	models.DB.Find(&foods)

	return c.Status(200).JSON(&fiber.Map{
		"data":    &foods,
		"success": true,
		"message": "food list",
	})
}

func RetrieveFood(c *fiber.Ctx) error {
	id := c.Params("id")

	var food models.FoodItem
	models.DB.First(&food, id)

	return c.Status(200).JSON(&fiber.Map{
		"data":    &food,
		"success": true,
		"message": "food detail",
	})
}

func InsertFood(c *fiber.Ctx) error {
	var food models.FoodItem

	// Validate user input
	if err := c.BodyParser(&food); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "invalid food data",
			"error":   err.Error(),
			"success": false,
		})
	}

	// Insert new food data
	if err := models.DB.Create(&food).Error; err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "cannot insert food",
			"error":   err.Error(),
			"success": false,
		})
	}

	// Return success response
	return c.Status(200).JSON(&fiber.Map{
		"message": "success insert food",
		"data":    "ID: " + strconv.Itoa(int(food.ID)),
		"success": true,
	})
}

func UpdateFood(c *fiber.Ctx) error {
	id := c.Params("id")
	var food models.FoodItem

	// Find food with given ID
	result := models.DB.First(&food, id)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "food not found",
			"error":   result.Error.Error(),
			"success": false,
		})
	}

	// Validate user input
	var updatedFood models.FoodItem
	if err := c.BodyParser(&updatedFood); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request parameter",
			"error":   err.Error(),
			"success": false,
		})
	}

	// Update food data
	food.Name = updatedFood.Name
	food.Price = updatedFood.Price
	if err := models.DB.Save(&food).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot update food data",
			"error":   err.Error(),
			"success": false,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success update food data",
		"success": true,
	})
}

func PatchFood(c *fiber.Ctx) error {
	id := c.Params("id")
	var food models.FoodItem

	// Find food with given ID
	result := models.DB.First(&food, id)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "food not found",
			"error":   result.Error.Error(),
			"success": false,
		})
	}

	var updatedFood map[string]interface{}
	if err := c.BodyParser(&updatedFood); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request parameter",
			"error":   result.Error.Error(),
			"success": false,
		})
	}

	// Update food data
	if err := models.DB.Model(&food).Updates(updatedFood).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to update food",
			"error":   result.Error.Error(),
			"success": false,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success update food data",
		"success": true,
	})
}
