package json_storage

import (
	"TaskTracker/internal/entities"
	"TaskTracker/pkg/file_functions"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var ErrTaskNotFound = errors.New("task not found")

type JsonStorage struct {
	basePath string
}

// New Create new JsonStorage
func New(basePath string) (*JsonStorage, error) {
	err := file_functions.CreateFileIfNotExists(basePath)
	if err != nil {
		return nil, err
	}

	return &JsonStorage{basePath}, nil
}

// loading task from json file and return slice of Tasks
func (js JsonStorage) loadTasksFromStorage() ([]*entities.Task, error) {
	// Crete a file if it does not exist
	if err := file_functions.CreateFileIfNotExists(js.basePath); err != nil {
		return nil, fmt.Errorf("cannot create file: %v", err)
	}

	// Read file data
	data, err := os.ReadFile(js.basePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %v", err)
	}

	// Unmarshal data to json array
	var tasks []*entities.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("cannot unmarshal data: %v", err)
	}

	return tasks, nil
}

func (js JsonStorage) loadTasksToStorage(tasks []*entities.Task) error {
	// Marshal json array
	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("cannot marshal data: %v", err)
	}

	if err := os.WriteFile(js.basePath, data, 0644); err != nil {
		return fmt.Errorf("cannot write file: %v", err)
	}

	return nil
}

func (js JsonStorage) isTaskInStorage(task *entities.Task) (bool, error) {
	tasks, err := js.loadTasksFromStorage()
	if err != nil {
		return false, err
	}

	for _, curTask := range tasks {
		if curTask.Id == task.Id {
			return true, nil
		}
	}

	return false, nil
}

func (js JsonStorage) AddNewTask(task *entities.Task) error {
	tasks, err := js.loadTasksFromStorage()
	if err != nil {
		return err
	}

	// Add new task to json array
	tasks = append(tasks, task)

	err = js.loadTasksToStorage(tasks)
	if err != nil {
		return err
	}

	return nil
}

func (js JsonStorage) NumOfTasks() (int, error) {
	tasks, err := js.loadTasksFromStorage()
	if err != nil {
		return -1, err
	}

	return len(tasks), nil
}

func (js JsonStorage) GetTasks() ([]*entities.Task, error) {
	tasks, err := js.loadTasksFromStorage()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (js JsonStorage) GetFreeId() (int, error) {
	tasks, err := js.loadTasksFromStorage()
	if err != nil {
		return -1, err
	}

	if len(tasks) == 0 {
		return 0, nil
	}

	return tasks[len(tasks)-1].Id + 1, nil
}

func (js JsonStorage) GetTaskById(id int) (*entities.Task, error) {
	storage, err := js.loadTasksFromStorage()
	if err != nil {
		return nil, err
	}

	for _, task := range storage {
		if task.Id == id {
			return task, nil
		}
	}

	return nil, nil
}

func (js JsonStorage) GetAllTasks() ([]*entities.Task, error) {
	tasks, err := js.loadTasksFromStorage()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (js JsonStorage) GetTasksByStatus(status entities.Status) ([]*entities.Task, error) {
	tasks, err := js.loadTasksFromStorage()
	if err != nil {
		return nil, err
	}

	completedTasks := make([]*entities.Task, 0)
	for _, task := range tasks {
		if task.Status == status {
			completedTasks = append(completedTasks, task)
		}
	}

	return completedTasks, nil
}

func (js JsonStorage) UpdateTaskStatus(id int, newStatus entities.Status) error {
	tasks, err := js.loadTasksFromStorage()
	if err != nil {
		return err
	}

	for _, curTask := range tasks {
		if curTask.Id == id {
			curTask.Status = newStatus
			return nil
		}
	}

	return ErrTaskNotFound
}
