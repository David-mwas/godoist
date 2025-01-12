package handler

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// Declare the Todo struct as you did in the original code
type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

// Initialize MongoDB connection
func init() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collection = client.Database(os.Getenv("MONGODB_DB")).Collection("todos")
}

// Handler is the main entry point of the application
func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	app := handler()
	app.ServeHTTP(w, r)
}

// Building the Fiber application
func handler() http.Handler {
	app := fiber.New()

	// Route for fetching all todos
	app.Get("/api/todo", func(c *fiber.Ctx) error {
		todos := []map[string]interface{}{}
		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var todo map[string]interface{}
			if err := cursor.Decode(&todo); err != nil {
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			}
			todos = append(todos, todo)
		}
		return c.JSON(todos)
	})

	// Route for adding a new todo
	app.Post("/api/todo", func(c *fiber.Ctx) error {
		var todo Todo
		if err := c.BodyParser(&todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "todo Body is required"})
		}

		// Insert the todo document into MongoDB
		insertResult, err := collection.InsertOne(context.Background(), todo)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		// Set the ID after the insert and return the response
		todo.ID = insertResult.InsertedID.(primitive.ObjectID)
		return c.Status(201).JSON(todo)
	})

	// Route for updating a todo
	app.Patch("/api/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid todo id"})
		}

		filter := bson.M{"_id": objectID}
		update := bson.M{"$set": bson.M{"completed": true}}
		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(200).JSON(fiber.Map{"success": true})
	})

	// Route for deleting a todo
	app.Delete("/api/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid todo id"})
		}

		filter := bson.M{"_id": objectID}
		_, err = collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(200).JSON(fiber.Map{"success": true})
	})

	return adaptor.FiberApp(app)
}
