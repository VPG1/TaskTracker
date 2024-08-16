package usecases

import (
	"TaskTracker/internal/entities"
	"TaskTracker/internal/usecases/mocks"
	"testing"
)

func TestAddTask(t *testing.T) {
	type args struct {
		name        string
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "base test",
			args: args{
				"task1",
				"description1",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := mocks.NewTaskStorage(t)

			storage.On("SaveTask", entities.CreateTask(0, tt.args.name, tt.args.description)).Return(nil)
			storage.On("GetFreeId").Return(0, nil)

			got, err := AddTask(storage, tt.args.name, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("AddTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}
