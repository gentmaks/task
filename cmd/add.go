//Package cmd
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/gentmaks/services"
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
}
