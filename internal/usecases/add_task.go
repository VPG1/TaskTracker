package usecases

import (
	"TaskTracker/internal/entities"
	"fmt"
)

type TaskStorage interface {
	SaveTask(e *entities.Task) error
	GetFreeId() (int, error)
}

func AddTask(storage TaskStorage, name string, description string) (int, error) {
	freeId, err := storage.GetFreeId()
	if err != nil {
		return -1, err
	}

	newTask := entities.CreateTask(freeId, name, description)

	if err := storage.SaveTask(newTask); err != nil {
		return -1, fmt.Errorf("cant save task to storage")
	}

	//log.Printf("Add task to storage: %s", newTask)

	return freeId, nil
}
