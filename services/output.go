//Package services
package services

import (
	"github.com/gentmaks/task/models"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"fmt"
)

func PrintTasks(taskList []models.Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID\t", "\t\tDescription\t\t", "\t\tState\t\t"})
	t.AppendSeparator()
	for _, task := range taskList {
		t.AppendRow([]interface{}{task.ID, task.Description, stateColor(task.State)})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
}

// stateColor func - Give a custom color output to the state of task
func stateColor(state string) string {
	switch state {
	case "In-progress":
		return fmt.Sprintf("🚀 %s%s%s\n", colorBlue, state, colorReset)
	case "Completed":
		return fmt.Sprintf("✅ %s%s%s\n", colorGreen, state, colorReset)
	case "Open":
		return fmt.Sprintf("📌 %s%s%s\n", colorGray, state, colorReset)
	}
	return state
}
