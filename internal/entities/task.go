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
