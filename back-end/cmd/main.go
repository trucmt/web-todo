package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/trucmt/web-todo/back-end/db"
	"github.com/trucmt/web-todo/back-end/handlers"
)

func main() {
	db.ConnectDB()

	app := fiber.New()

	app.Get("/todos", handlers.GetAllTodos)
	app.Post("/todos", handlers.CreateTodo)
	// Thêm các route khác

	log.Fatal(app.Listen(":4000"))
}
