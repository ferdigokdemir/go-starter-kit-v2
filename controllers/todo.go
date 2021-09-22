package controllers

import (
	"go_starter_kit/config"
	"go_starter_kit/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetTodos : get all todos
func GetTodos(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection("todos")

	// find all todos
	cursor, err := todoCollection.Find(c.Context(), bson.D{})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot find todos",
			"error":   err,
		})
	}

	// iterate through cursor
	todos := []*models.Todo{}
	for cursor.Next(c.Context()) {
		todo := &models.Todo{}
		cursor.Decode(todo)
		todos = append(todos, todo)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}

// CreateTodo : Create a todo
func CreateTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection("todos")

	// var data Request
	data := new(models.Todo)
	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse body",
			"error":   err,
		})
	}

	// create todo
	todo := &models.Todo{
		Title:     data.Title,
		Completed: data.Completed,
	}

	// insert todo
	_, err = todoCollection.InsertOne(c.Context(), todo)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot create todo",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

// GetTodo : get a single todo
// PARAM: id
func GetTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection("todos")

	// get parameter value
	paramID := c.Params("id")

	// convert parameterID to objectId
	id, err := primitive.ObjectIDFromHex(paramID)

	// if error while parsing paramID
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
			"error":   err,
		})
	}

	// find todo and return

	todo := &models.Todo{}

	query := bson.D{{Key: "_id", Value: id}}

	err = todoCollection.FindOne(c.Context(), query).Decode(todo)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Todo Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

// UpdateTodo : Update a todo
// PARAM: id
func UpdateTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection("todos")

	// get parameter value
	paramID := c.Params("id")

	// convert parameterID to objectId
	id, err := primitive.ObjectIDFromHex(paramID)

	// if error while parsing paramID
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
			"error":   err,
		})
	}

	// var data Request
	data := new(models.Todo)
	err = c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse body",
			"error":   err,
		})
	}

	// update todo
	todo := &models.Todo{
		Title:     data.Title,
		Completed: data.Completed,
	}

	// update todo
	_, err = todoCollection.UpdateOne(c.Context(), bson.D{{Key: "_id", Value: id}}, bson.D{{Key: "$set", Value: todo}})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update todo",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

// DeleteTodo : Delete a todo
// PARAM: id
func DeleteTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection("todos")

	// get parameter value
	paramID := c.Params("id")

	// convert parameterID to objectId
	id, err := primitive.ObjectIDFromHex(paramID)

	// if error while parsing paramID
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
			"error":   err,
		})
	}

	// delete todo
	_, err = todoCollection.DeleteOne(c.Context(), bson.D{{Key: "_id", Value: id}})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete todo",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id": id,
		},
	})
}
