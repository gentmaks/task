//Package cmd
package cmd

import (
	"github.com/spf13/cobra"
	"log"
)



var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Go-based CLI tool to take care of your todo list",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
