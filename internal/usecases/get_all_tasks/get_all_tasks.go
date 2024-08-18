package get_all_tasks

import "TaskTracker/internal/entities"

type TaskStorage interface {
	GetAllTasks() ([]*entities.Task, error)
}

func GetAllTasks(store TaskStorage) ([]*entities.Task, error) {
	return store.GetAllTasks()
}
