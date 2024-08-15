package main

import (
	"TaskTracker/internal/storages/json_storage"
	"TaskTracker/internal/usecase"
	"fmt"
	"log"
)

func main() {
	storage, err := json_storage.New("storage.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = usecase.AddTask(storage, "Task1", "desc")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Hello World")
}
