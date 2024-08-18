package entities

import (
	"testing"
	"time"
)

// function do deep equal of Task structs without checking CreatedAt
func taskIsEqual(left *Task, right *Task) bool {
	return left.Id == right.Id && left.Name == right.Name && left.Description == right.Description &&
		left.Status == right.Status
}

func TestCreateTask(t *testing.T) {
	type args struct {
		id          int
		name        string
		description string
	}
	tests := []struct {
		name string
		args args
		want *Task
	}{
		{
			name: "Task Creating test",
			args: args{
				id:          1,
				name:        "task1",
				description: "task1 description",
			},
			want: &Task{1, "task1", "task1 description", NotDone, time.Time{}, time.Time{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTask(tt.args.id, tt.args.name, tt.args.description); !taskIsEqual(got, tt.want) {
				t.Errorf("CreateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
