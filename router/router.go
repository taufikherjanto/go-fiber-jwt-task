package router

import (
	"go-fiber-jwt-task/controller"
	"go-fiber-jwt-task/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Authentication
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)

	// Tasks
	task := api.Group("/tasks")
	task.Get("/", controller.GetTasks)
	task.Get("/:id", controller.GetTask)

	// Task with authorization
	task.Use(middleware.JWTAuthorization)
	task.Post("/", controller.CreateTask)
	task.Patch("/:id", controller.UpdateTask)
	task.Patch("/:id/done", controller.UpdateDoneTask)
	task.Delete("/:id", controller.DeleteTask)
}
