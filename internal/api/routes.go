package api

import (
	"github.com/gofiber/fiber/v2"
	"simple-service/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/tasks", handlers.CreateTask)
	app.Get("/tasks", handlers.GetAllTasks)
	app.Get("/tasks/:id", handlers.GetTaskByID)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)
}
