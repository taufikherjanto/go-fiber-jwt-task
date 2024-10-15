package controller

import (
	"go-fiber-jwt-task/database"
	"go-fiber-jwt-task/model"
	"go-fiber-jwt-task/utils"

	"github.com/gofiber/fiber/v2"
)

// The request Dto for both register and login
type authenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var req authenticationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := model.User{
		Email:        req.Email,
		PasswordHash: utils.GeneratePassword(req.Password),
	}

	// Check if user exists
	existingUser := database.DB.Where("email = ?", req.Email).First(&user)
	if existingUser.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	response := database.DB.Create(&user)
	if response.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": response.Error.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "user created",
	})
}

func Login(c *fiber.Ctx) error {
	var req authenticationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Check if user exists
	var user model.User
	response := database.DB.Where("email = ?", req.Email).First(&user)
	if response.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if !utils.ComparePassword(user.PasswordHash, req.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password is incorrect",
		})
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"token": token,
	})
}
