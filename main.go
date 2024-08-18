package main

import (
	"TaskTracker/internal/storages/json_storage"
	"TaskTracker/internal/usecases/add_task"
	"fmt"
	"log"
)

func main() {
	storage, err := json_storage.New("storage.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = add_task.AddTask(storage, "Task1", "desc")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Hello World")
}
