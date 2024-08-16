package file_functions

import (
	"fmt"
	"log"
	"os"
)

func CheckIsFileExist(path string) bool {
	// Checking is file exist
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	return true
}

func CreateFileIfNotExists(path string) error {
	if !CheckIsFileExist(path) {
		file, err := os.Create(path)
		fmt.Println(file.Stat())
		if err != nil {
			log.Fatal("cannot create file: ", err)
			return err
		}

		_, err = file.Write([]byte("[]"))
		if err != nil {
			return err
		}
	}

	return nil
}
