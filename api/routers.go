package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trucmt/web-todo/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Trang chủ
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// API routes cho todos
	todoRoutes := app.Group("/api/todos")
	todoRoutes.Get("/", handlers.GetAllTodos)
	todoRoutes.Get("/:id", handlers.GetTodo)
	todoRoutes.Post("/", handlers.CreateTodo)
	// todoRoutes.Put("/:id", handlers.UpdateTodo)
	// todoRoutes.Delete("/:id", handlers.DeleteTodo)

	// Thêm các routes khác nếu cần
	// Ví dụ: routes cho users, authentication, etc.
}
