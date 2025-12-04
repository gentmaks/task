//Package cmd
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/gentmaks/task/services"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets the list of tasks by asserting a clean todos.txt file",
	Run: reset,
}

func init() {
	rootCmd.AddCommand(resetCmd)
}

func reset(cmd *cobra.Command, args []string) {
	services.OverwriteFile()	
}
