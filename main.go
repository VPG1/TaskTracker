package main

import (
	"TaskTracker/internal/entities"
	"TaskTracker/internal/storages/json_storage"
	"TaskTracker/internal/usecases/get_tasks_by_status"
	"fmt"
	"log"
)

func main() {
	storage, err := json_storage.New("storage.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	v, err := get_tasks_by_status.GetTasksByStatus(storage, entities.NotDone+5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)
}
