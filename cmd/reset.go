//Package cmd
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
	fmt.Println("reset called")	
}
