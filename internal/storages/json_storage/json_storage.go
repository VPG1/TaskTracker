package json_storage

import (
	"TaskTracker/internal/entities"
	"TaskTracker/pkg/files_function"
	"encoding/json"
	"fmt"
	"os"
)

type JsonStorage struct {
	basePath string
}

func New(basePath string) (*JsonStorage, error) {
	err := files_function.CreateFileIfNotExists(basePath)
	if err != nil {
		return nil, err
	}

	return &JsonStorage{basePath}, nil
}

func (js JsonStorage) SaveTask(task *entities.Task) error {
	// Crete a file if it does not exist
	if err := files_function.CreateFileIfNotExists(js.basePath); err != nil {
		return fmt.Errorf("cannot create file: %v", err)
	}

	// Read file data
	data, err := os.ReadFile(js.basePath)
	if err != nil {
		return fmt.Errorf("cannot read file: %v", err)
	}

	// Unmarshal data to json array
	var tasks []entities.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return fmt.Errorf("cannot unmarshal data: %v", err)
	}

	// Add new task to json array
	tasks = append(tasks, *task)

	// Marshal json array
	data, err = json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("cannot marshal data: %v", err)
	}

	if err := os.WriteFile(js.basePath, data, 0644); err != nil {
		return fmt.Errorf("cannot write file: %v", err)
	}

	return nil
}

func (js JsonStorage) NumOfTasks() (int, error) {
	// Crete a file if it does not exist
	if err := files_function.CreateFileIfNotExists(js.basePath); err != nil {
		return -1, fmt.Errorf("cannot create file: %v", err)
	}

	// Read file data
	data, err := os.ReadFile(js.basePath)
	if err != nil {
		return -1, fmt.Errorf("cannot read file: %v", err)
	}

	// Unmarshal data to json array
	var tasks []entities.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return -1, fmt.Errorf("cannot unmarshal data: %v", err)
	}

	return len(tasks), nil
}
