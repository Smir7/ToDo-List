package repo

import (
	"simple-service/internal/models"
)

var tasks = make(map[int]models.Task)
var currentID = 1

func CreateTask(task models.Task) models.Task {
	task.ID = currentID
	tasks[currentID] = task
	currentID++
	return task
}

func GetAllTasks() []models.Task {
	var taskList []models.Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return taskList
}

func GetTaskByID(id int) (models.Task, bool) {
	task, exists := tasks[id]
	return task, exists
}

func UpdateTask(id int, updatedTask models.Task) bool {
	if _, exists := tasks[id]; exists {
		updatedTask.ID = id
		tasks[id] = updatedTask
		return true
	}
	return false
}

func DeleteTask(id int) bool {
	if _, exists := tasks[id]; exists {
		delete(tasks, id)
		return true
	}
	return false
}
