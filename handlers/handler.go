package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trucmt/web-todo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var todos []models.Todo

func GetAllTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
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
		{ID: primitive.NewObjectID(), Title: "Todo 1", Description: "Body 1", Completed: false},
		{ID: primitive.NewObjectID(), Title: "Todo 2", Description: "Body 2", Completed: false},
		{ID: primitive.NewObjectID(), Title: "Todo 3", Description: "Body 3", Completed: false},
	}
}
