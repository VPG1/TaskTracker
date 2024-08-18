package main

import (
	"TaskTracker/internal/storages/json_storage"
	"TaskTracker/internal/usecases/get_task_by_id"
	"fmt"
	"log"
)

func main() {
	storage, err := json_storage.New("storage.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	v, err := get_task_by_id.GetTaskById(storage, 5)
	if err != nil {
		return
	}

	fmt.Println(v)
}
