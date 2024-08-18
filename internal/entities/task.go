package entities

import "time"

type Status int

const (
	NotDone = iota
	InProgress
	Done
)

type Task struct {
	Id          int
	Name        string
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateTask(id int, name string, description string) *Task {
	t := time.Now()
	return &Task{Id: id, Name: name, Description: description,
		Status: NotDone, CreatedAt: t, UpdatedAt: t}
}
