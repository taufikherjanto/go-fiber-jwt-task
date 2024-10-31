package router

import (
	"go-fiber-jwt-task/controller"
	"go-fiber-jwt-task/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes menginisialisasi semua rute API.
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api") // Grup API utama

	// Rute Autentikasi
	auth := api.Group("/auth")                                                 // Grup untuk rute terkait autentikasi
	auth.Post("/login", controller.Login)                                      // Rute untuk login pengguna
	auth.Post("/register", controller.Register)                                // Rute untuk pendaftaran pengguna
	auth.Get("/userinfo", middleware.JWTAuthorization, controller.GetUserInfo) // Rute info pengguna yang dilindungi

	// Rute Tugas
	task := api.Group("/tasks")          // Grup untuk rute terkait tugas
	task.Get("/", controller.GetTasks)   // Mengambil semua tugas
	task.Get("/:id", controller.GetTask) // Mengambil tugas tertentu

	// Rute Tugas yang Dilindungi
	task.Use(middleware.JWTAuthorization)              // Terapkan middleware JWT untuk melindungi rute tugas
	task.Post("/", controller.CreateTask)              // Membuat tugas baru
	task.Patch("/:id", controller.UpdateTask)          // Memperbarui tugas yang ada
	task.Patch("/:id/done", controller.UpdateDoneTask) // Menandai tugas sebagai selesai
	task.Delete("/:id", controller.DeleteTask)         // Menghapus tugas tertentu
}
