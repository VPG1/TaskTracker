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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := add_task.AddTask(JsonStorage, taskName, taskDescription)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&taskName, "name", "", "task name")
	err := addCmd.MarkFlagRequired("name")
	if err != nil {
		return
	}
	addCmd.Flags().StringVar(&taskDescription, "description", "", "task description")
	err = addCmd.MarkFlagRequired("description")
	if err != nil {
		return
	}
}
