//Package services
package services

import (
	"os"
	"log"
)

const fileName string = "todos.txt"

func CreateFile() *os.File{
	file, err := os.Create(fileName)	
	if err != nil {
		log.Fatal("Failed to create the file")
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatal("Failed to close the file", err)
		}
	}(file)
	return file
}

func OpenFile() *os.File{
	file, err := os.OpenFile(fileName, os.O_CREATE | os.O_WRONLY | os.O_APPEND, os.ModePerm)
	if err != nil {
		return CreateFile()
	}
	return file
}

func ReadFile() []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Could not read from the file:", err)
	}
	return data
}

func WriteFile(data []byte) {
	file := OpenFile()
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatal("Could not close the file:", err)
		}
	}(file)
	_, err := file.WriteString(string(data))
	if err != nil {
		log.Fatal("Could not write to the todos file:", err)
	}
}

func OverwriteFile() {
	file := OpenFile()
	err := file.Truncate(0)
	if err != nil {
		log.Fatal("Error truncating file", err)
	}
}
