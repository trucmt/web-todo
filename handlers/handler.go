package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trucmt/web-todo/models"
)

var todos []models.Todo

func GetAllTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, todo := range todos {
		if todo.ID == id {
			return c.JSON(todo)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func CreateTodo(c *fiber.Ctx) error {
	var newTodo models.Todo
	if err := c.BodyParser(&newTodo); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	todos = append(todos, newTodo)
	return c.JSON(newTodo)
}

// Khởi tạo dữ liệu mẫu
func InitTodos() {
	todos = []models.Todo{
		{ID: "1", Title: "Todo 1", Body: "Body 1", Completed: false},
		{ID: "2", Title: "Todo 2", Body: "Body 2", Completed: false},
		{ID: "3", Title: "Todo 3", Body: "Body 3", Completed: false},
	}
}
