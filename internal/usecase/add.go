package usecase

import (
	"TaskTracker/internal/entities"
	"fmt"
	"log"
)

type TaskStorage interface {
	SaveTask(e *entities.Task) error
	NumOfTasks() (int, error)
}

func AddTask(storage TaskStorage, name string, description string) (int, error) {
	numOfTasks, err := storage.NumOfTasks()

	if err != nil {
		return -1, fmt.Errorf("Cant calculate new task id")
	}

	newTask := entities.CreateTask(numOfTasks, name, description)

	if err := storage.SaveTask(newTask); err != nil {
		return -1, fmt.Errorf("Cant save task to storage")
	}

	log.Printf("Add task to storage: %s", newTask)

	return numOfTasks, nil
}
