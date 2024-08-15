package entities

import "time"

const (
	NotDone = iota
	InProgress
	Done
)

type Task struct {
	id          int
	name        string
	description string
	status      int
	createdAt   time.Time
	completeAt  time.Time
}

func CreateTask(id int, name string, description string) *Task {
	return &Task{id: id, name: name, description: description, status: NotDone, createdAt: time.Now(), completeAt: time.Time{}}
}
