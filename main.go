package main

import (
	"TaskTracker/internal/storages/json_storage"
	"TaskTracker/internal/usecases/delete_task"
	"fmt"
	"log"
)

func main() {
	storage, err := json_storage.New("storage.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = delete_task.DeleteTask(storage, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
}
