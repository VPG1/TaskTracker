package cmd

import (
	"TaskTracker/internal/task_tracker_errors"
	"TaskTracker/internal/usecases/delete_task"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <task_id>",
	Args:  cobra.ExactArgs(1),
	Short: "Command to delete a task",
	Long: `This command will delete a task from the task tracker by task ID.
Command has one required argument with a task ID.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("argument is not a number")
		}

		err = delete_task.DeleteTask(JsonStorage, id)
		if errors.Is(err, task_tracker_errors.ErrTaskNotFound) {
			return fmt.Errorf("task with id %v not found", id)
		} else if err != nil {
			return err
		}

		fmt.Printf("Task with id %v deleted", id)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
