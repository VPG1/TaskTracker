package usecase

import (
	"TaskTracker/internal/entities"
	"log"
)

type TaskStorage interface {
	SaveTask(e *entities.Task) error
	NumOfTasks() (int, error)
}

func AddTask(storage TaskStorage, name string, description string) {
	numOfTasks, err := storage.NumOfTasks()

	if err != nil {
		log.Fatal("Cant calculate new task id")
		return
	}

	newTask := entities.CreateTask(numOfTasks, name, description)

	if err := storage.SaveTask(newTask); err != nil {
		log.Fatal("Cant save task to storage")
		return
	}

	log.Printf("Add task to storage: %s", newTask)
}
