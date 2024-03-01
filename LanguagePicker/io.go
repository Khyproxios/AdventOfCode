package main

import (
	"fmt"
	"os"
)

func verifyStore(filepath string) {
	_, error := os.Stat(filepath)

	if error != nil {
		_, error := os.Create(filepath)

		if error != nil {
			fmt.Println("Error with creating the file:", error)

			return
		}
	}
}

func loadData(filename string) ([]byte, error) {
	data, error := os.ReadFile(filename)

	if error != nil {
		return nil, error
	}

	return data, error
}

func saveData(file string, data []byte) error {
	// TODO: Verify the umask used here works
	return os.WriteFile(file, data, 0777)
}
