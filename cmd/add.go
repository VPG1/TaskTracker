package cmd

import (
	"TaskTracker/internal/usecases/add_task"
	"fmt"
	"github.com/spf13/cobra"
)

var taskName string
var taskDescription string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Command to add a new task",
	Long: `This command will add a new task to the task tracker.
The command has two required flags --name and --description
to set the name and description of the new task`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := add_task.AddTask(JsonStorage, taskName, taskDescription)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Added Task: ", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&taskName, "name", "n", "", "task name")
	err := addCmd.MarkFlagRequired("name")
	if err != nil {
		return
	}
	addCmd.Flags().StringVarP(&taskDescription, "description", "d", "", "task description")
	err = addCmd.MarkFlagRequired("description")
	if err != nil {
		return
	}
}
