package controller

import (
	"go-fiber-jwt-task/database"
	"go-fiber-jwt-task/model"

	"github.com/gofiber/fiber/v2"
)

type taskRequest struct {
	Title       string `gorm:"not null" json:"Title"`
	Description string `json:"description"`
}

func GetTasks(c *fiber.Ctx) error {
	var tasks []model.Task
	database.DB.Find(&tasks)
	return c.JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task model.Task
	database.DB.Find(&task, id)
	return c.JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	var task taskRequest
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	taskModel := model.Task{
		Title:       task.Title,
		Description: task.Description,
	}

	response := database.DB.Create(&taskModel)
	if response.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": response.Error.Error(),
		})
	}

	return c.JSON(task)
}
