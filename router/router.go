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
	tasks := api.Group("/tasks")
	tasks.Post("/", middleware.JWTAuthorization, controller.CreateTask)
}
