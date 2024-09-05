package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/trucmt/web-todo/api"
	"github.com/trucmt/web-todo/handlers"
)

func main() {
	app := fiber.New()

	handlers.InitTodos()
	api.SetupRoutes(app)

	log.Fatal(app.Listen(":4000"))
}
