//Package cmd
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/gentmaks/task/services"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete/done",
	Run: do,
}

func init() {
	rootCmd.AddCommand(doCmd)
}

func do(cmd *cobra.Command, args[]string) {
	services.EditTask(args)
}
