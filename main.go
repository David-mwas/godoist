package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello Worlds you")

	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	todos := []Todo{}

	// middleware
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello, Welcome to go todolist"})
	})

	// get all todos
	app.Get(("/api/todo"), func(c *fiber.Ctx) error {

		return c.Status(200).JSON(todos)
	})

	// create todo
	app.Post("/api/todo", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)

	})

	// update
	app.Patch("/api/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])
			}

		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	// delete
	app.Delete("/api/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": "Todo deleted"})
			}

		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})

	})
	PORT := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + PORT))
}
