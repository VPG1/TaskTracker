package get_tasks_by_status

import (
	"TaskTracker/internal/entities"
	"errors"
)

type TaskStorage interface {
	GetTasksByStatus(status entities.Status) ([]*entities.Task, error)
}

func GetTasksByStatus(storage TaskStorage, status entities.Status) ([]*entities.Task, error) {
	if status < 0 || status > 2 {
		return nil, errors.New("invalid status")
	}

	return storage.GetTasksByStatus(status)
}
