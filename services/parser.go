//Package services
package services

import (
	"github.com/gentmaks/task/models"
	"gopkg.in/yaml.v2"
	"log"
	"strconv"
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

func GenerateID() int {
	taskList := ReadData()
	if taskList == nil {
		return 1
	}
	return taskList[len(taskList) - 1].ID + 1 
}

func EditTask(args []string) {
	if len(args) != 2 {
		log.Fatalln("Usage: do <ID of task> <[Completed, In-Progress, Open]>")
	}
	taskList := ReadData()
	index, err := strconv.Atoi(args[0])
	if err != nil || index < 0 || index > len(taskList) {
		log.Fatalln("Invalid Id detected")
	}
	if args[1] != "Completed" && args[1] != "In-Progress" && args[1] != "Open" {
		log.Fatalln("Invalid state passed")
	}
	OverwriteFile()
	WriteData(taskList)
}
