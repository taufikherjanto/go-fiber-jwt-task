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

type updateDoneRequest struct {
	Done bool `json:"done"`
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
	if task.Title == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

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

func UpdateTask(c *fiber.Ctx) error {
	var task taskRequest
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id := c.Params("id")
	var dataTask model.Task
	database.DB.Find(&dataTask, id)
	if dataTask.Title == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	database.DB.Model(&dataTask).Updates(task)
	return c.JSON(task)
}

func UpdateDoneTask(c *fiber.Ctx) error {
	var task updateDoneRequest
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id := c.Params("id")
	var dataTask model.Task
	database.DB.Find(&dataTask, id)
	if dataTask.Title == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	database.DB.Model(&dataTask).Updates(task)
	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task model.Task
	database.DB.Find(&task, id)
	if task.Title == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	//run delete data
	database.DB.Delete(&task)
	return c.SendString("Task deleted")
}
