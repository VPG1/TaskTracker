package update_task

import "TaskTracker/internal/entities"

type Storage interface {
	GetTaskById(id int) (*entities.Task, error)
	UpdateTask(id int, task *entities.Task) error
}

func UpdateTask(storage Storage, id int, name, description string) error {
	task, err := storage.GetTaskById(id)
	if err != nil {
		return err
	}

	task.Name = name
	task.Description = description

	err = storage.UpdateTask(id, task)
	if err != nil {
		return err
	}

	return nil
}
