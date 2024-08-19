package mark_task

import (
	"TaskTracker/internal/entities"
	"TaskTracker/internal/task_tracker_errors"
)

type TaskStorage interface {
	UpdateTaskStatus(id int, task entities.Status) error
}

func MarkTask(storage TaskStorage, id int, status entities.Status) error {
	if status < 0 || status > 2 {
		return task_tracker_errors.ErrInvalidStatus
	}

	err := storage.UpdateTaskStatus(id, status)
	if err != nil {
		return err
	}

	return nil
}
