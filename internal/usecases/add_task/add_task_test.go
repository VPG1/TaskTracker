package add_task

import (
	"TaskTracker/internal/usecases/add_task/mocks"
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAddTask(t *testing.T) {
	t.Run("base test", func(t *testing.T) {
		storage := mocks.NewTaskStorage(t)
		storage.On("GetFreeId").Return(0, nil)
		storage.On("SaveTask", mock.Anything).Return(nil)

		_, err := AddTask(storage, "task1", "description")
		if err != nil {
			t.Errorf("AddTask() error = %v, wantErr %v", err, nil)
		}
	})

	t.Run("GetFreeId return error", func(t *testing.T) {
		storage := mocks.NewTaskStorage(t)
		storage.On("GetFreeId").Return(-1, fmt.Errorf("GetFreeId error"))

		_, err := AddTask(storage, "task2", "description")
		if err == nil {
			t.Errorf("AddTask() error = %v, wantErr %v", nil, fmt.Errorf("GetFreeId error"))
		}
	})

	t.Run("SaveTask returns error", func(t *testing.T) {
		storage := mocks.NewTaskStorage(t)
		storage.On("GetFreeId").Return(0, nil)
		storage.On("SaveTask", mock.Anything).Return(fmt.Errorf("SaveTask error"))

		_, err := AddTask(storage, "task1", "description")
		if err == nil {
			t.Errorf("AddTask() error = %v, wantErr %v", nil, fmt.Errorf("SaveTask error"))
		}
	})
}
