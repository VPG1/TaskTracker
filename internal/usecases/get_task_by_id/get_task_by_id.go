package get_task_by_id

import "TaskTracker/internal/entities"

type TaskStorage interface {
	GetTaskById(id int) (*entities.Task, error)
}

func GetTaskById(storage TaskStorage, id int) (*entities.Task, error) {
	return storage.GetTaskById(id)
}
