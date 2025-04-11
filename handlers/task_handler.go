package handlers

import (
	"github.com/gofiber/fiber/v2"
	"simple-service/internal/models"
	"simple-service/internal/repo"
	"strconv"
)

func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	createdTask := repo.CreateTask(task)
	return c.Status(fiber.StatusCreated).JSON(createdTask)
}

func GetAllTasks(c *fiber.Ctx) error {
	tasks := repo.GetAllTasks()
	return c.Status(fiber.StatusOK).JSON(tasks)
}

func GetTaskByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	task, exists := repo.GetTaskByID(id)
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.Status(fiber.StatusOK).JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updatedTask models.Task
	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if !repo.UpdateTask(id, updatedTask) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.Status(fiber.StatusOK).JSON(updatedTask)
}

func DeleteTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if !repo.DeleteTask(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.Status(fiber.StatusNoContent).SendString("")
}
