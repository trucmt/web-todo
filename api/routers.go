package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trucmt/web-todo/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/todos", handlers.GetAllTodos)
	app.Get("/api/todos/:id", handlers.GetTodo)
	app.Post("/api/todos", handlers.CreateTodo)
}
