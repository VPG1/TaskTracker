package main

import (
	"TaskTracker/internal/entities"
	"TaskTracker/internal/storages/json_storage"
	"TaskTracker/internal/usecases/get_tasks_by_status"
	"TaskTracker/internal/usecases/mark_task"
	"fmt"
	"log"
)

func main() {
	storage, err := json_storage.New("storage.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	v, err := get_tasks_by_status.GetTasksByStatus(storage, entities.InProgress)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, task := range v {
		fmt.Println(*task)
	}

	err = mark_task.MarkTask(storage, 1, entities.InProgress)
	if err != nil {
		fmt.Println(err)
		return
	}

	v, err = get_tasks_by_status.GetTasksByStatus(storage, entities.InProgress+5)
	if err != nil {
		fmt.Println(err)
	}
	for _, task := range v {
		fmt.Println(*task)
	}
}
