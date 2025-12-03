//Package services
package services

import (
	"github.com/gentmaks/task/models"
	"gopkg.in/yaml.v2"
	"log"
)

func ParseData(taskList []models.Task) []byte {
	yamlData, err := yaml.Marshal(taskList)
	if err != nil {
		log.Fatal("Could not get the data")
	}
	return yamlData
}

func WriteData(taskList []models.Task) {
	data := ParseData(taskList)
	WriteFile(data)
} 

func ReadData() []models.Task {
	var taskList []models.Task
	data := ReadFile()
	err := yaml.Unmarshal(data, &taskList) 
	if err != nil {
		log.Fatal("Could not read the data")
	}
	return taskList
}

func GenerateID(taskList []models.Task) int {
	if len(taskList) == 0 {
		return 1
	}
	return taskList[len(taskList) - 1].ID + 1 
}



