package add_task

import (
	"TaskTracker/internal/entities"
)

type TaskStorage interface {
	SaveTask(*entities.Task) error
	GetFreeId() (int, error)
}

func AddTask(storage TaskStorage, name string, description string) (int, error) {
	freeId, err := storage.GetFreeId()
	if err != nil {
		return -1, err
	}

	newTask := entities.CreateTask(freeId, name, description)

	if err := storage.SaveTask(newTask); err != nil {
		return -1, err
	}

	return freeId, nil
}
