//Package cmd
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/gentmaks/task/services"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
	fmt.Println("list called")
	taskList := services.ReadData()
	services.PrintTasks(taskList)
}
