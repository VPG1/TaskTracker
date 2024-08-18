package mark_task

import (
	"TaskTracker/internal/entities"
	"TaskTracker/internal/usecases/usecase_errors"
)

type TaskStorage interface {
	UpdateTaskStatus(id int, task entities.Status) error
}

func MarkTask(storage TaskStorage, id int, status entities.Status) error {
	if status < 0 || status > 2 {
		return usecase_errors.ErrInvalidStatus
	}

	return nil
}
