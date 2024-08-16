package json_storage

import (
	"TaskTracker/internal/entities"
	"TaskTracker/pkg/file_functions"
	"encoding/json"
	"fmt"
	"os"
)

type JsonStorage struct {
	basePath string
}

func New(basePath string) (*JsonStorage, error) {
	err := file_functions.CreateFileIfNotExists(basePath)
	if err != nil {
		return nil, err
	}

	return &JsonStorage{basePath}, nil
}

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

func (js JsonStorage) SaveTask(task *entities.Task) error {
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
