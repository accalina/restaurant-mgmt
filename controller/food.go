package controller

import (
	"github.com/accalina/restaurant-mgmt/configuration"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/gofiber/fiber/v2"
)

func ListFood(c *fiber.Ctx) error {
	var foods []model.FoodItem
	configuration.DB.Find(&foods)

	return c.Status(200).JSON(model.GeneralResponse{
		Success: true,
		Message: "food list",
		Data:    &foods,
	})
}

func RetrieveFood(c *fiber.Ctx) error {
	id := c.Params("id")
	var food model.FoodItem
	configuration.DB.First(&food, id)

	return c.Status(200).JSON(model.GeneralResponse{
		Success: true,
		Message: "food detail",
		Data:    &food,
	})
}

func InsertFood(c *fiber.Ctx) error {
	var food model.FoodItem

	// Validate user input
	if err := c.BodyParser(&food); err != nil {
		return c.Status(400).JSON(model.GeneralResponse{
			Success: false,
			Message: "invalid food data",
			Data:    err,
		})
	}

	// Insert new food data
	if err := configuration.DB.Create(&food).Error; err != nil {
		return c.Status(400).JSON(model.GeneralResponse{
			Success: false,
			Message: "cannot insert food",
			Data:    err,
		})
	}

	// Return success response
	return c.Status(200).JSON(model.GeneralResponse{
		Success: true,
		Message: "success insert food",
		Data:    &food,
	})
}

func UpdateFood(c *fiber.Ctx) error {
	id := c.Params("id")
	var food model.FoodItem

	// Find food with given ID
	result := configuration.DB.First(&food, id)
	if result.Error != nil {
		return c.Status(400).JSON(model.GeneralResponse{
			Success: false,
			Message: "food not found",
			Data:    result.Error,
		})
	}

	// Validate user input
	var updatedFood model.FoodItem
	if err := c.BodyParser(&updatedFood); err != nil {
		return c.Status(400).JSON(model.GeneralResponse{
			Success: false,
			Message: "Invalid request parameter",
			Data:    err,
		})
	}

	// Update food data
	food.Name = updatedFood.Name
	food.Price = updatedFood.Price
	if err := configuration.DB.Save(&food).Error; err != nil {
		return c.Status(500).JSON(model.GeneralResponse{
			Success: false,
			Message: "Cannot update food data",
			Data:    err,
		})
	}

	return c.Status(200).JSON(model.GeneralResponse{
		Success: true,
		Message: "success update food",
		Data:    &food,
	})
}

func PatchFood(c *fiber.Ctx) error {
	id := c.Params("id")
	var food model.FoodItem

	// Find food with given ID
	result := configuration.DB.First(&food, id)
	if result.Error != nil {
		return c.Status(400).JSON(model.GeneralResponse{
			Success: false,
			Message: "food not found",
			Data:    result.Error,
		})
	}

	var updatedFood map[string]interface{}
	if err := c.BodyParser(&updatedFood); err != nil {
		return c.Status(400).JSON(model.GeneralResponse{
			Success: false,
			Message: "Invalid request parameter",
			Data:    result.Error,
		})
	}

	// Update food data
	if err := configuration.DB.Model(&food).Updates(updatedFood).Error; err != nil {
		return c.Status(400).JSON(model.GeneralResponse{
			Success: false,
			Message: "Failed to update food",
			Data:    result.Error,
		})
	}

	return c.Status(200).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success update food data",
		Data:    &food,
	})
}
