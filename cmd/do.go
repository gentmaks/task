//Package cmd
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
	fmt.Println("do called")
}
