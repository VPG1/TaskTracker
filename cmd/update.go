package cmd

import (
	"TaskTracker/internal/usecases/get_task_by_id"
	"TaskTracker/internal/usecases/update_task"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var newTaskName string
var newTaskDescription string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <task_id>",
	Args:  cobra.ExactArgs(1),
	Short: "Command to update a task",
	Long: `This command is used to update a task name and description by id.
You can chose which task you modify. 
Using flags --name and --description specify new task name and description.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("task_id must be integer")
		}

		if newTaskName == "" {
			task, err := get_task_by_id.GetTaskById(JsonStorage, id)
			if err != nil {
				return err
			}
			newTaskName = task.Name
		}
		if newTaskDescription == "" {
			task, err := get_task_by_id.GetTaskById(JsonStorage, id)
			if err != nil {
				return err
			}
			newTaskDescription = task.Description
		}

		err = update_task.UpdateTask(JsonStorage, id, newTaskName, newTaskDescription)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&newTaskName, "name", "n", "", "New task name")
	updateCmd.Flags().StringVarP(&newTaskDescription, "description", "d", "", "New task description")
}
