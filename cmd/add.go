//Package cmd
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/gentmaks/task/services"
	"github.com/gentmaks/task/models"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task on the list of tasks to be done",
	Run: add,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(cmd *cobra.Command, args []string) {
	fmt.Println("add called")
	if len(args) == 0 {
		fmt.Println("No args added for task")
		return
	}
	var taskList []models.Task
	for _, taskName := range(args) {
		id := services.GenerateID(taskList)
		task := models.Task{ID: id, Description: taskName, State: "Open"}
		taskList = append(taskList, task)
	}
	services.WriteData(taskList)
	fmt.Println("The task has been successfully added.")
}
