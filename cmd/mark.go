package cmd

import (
	"TaskTracker/internal/entities"
	"TaskTracker/internal/usecases/mark_task"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// markCmd represents the mark command
var markCmd = &cobra.Command{
	Use:   "mark <task_status> <task_id>",
	Args:  cobra.ExactArgs(2),
	Short: "The command change task status",
	Long: `This command change task status.
Command has two required arguments: <task_status> <task_id>.
You can select new status and id of the task you want to change by entering this two arguments.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("task_id must be integer")
		}

		switch args[0] {
		case "todo":
			err := mark_task.MarkTask(JsonStorage, id, entities.NotDone)
			if err != nil {
				return err
			}
		case "in-progress":
			err := mark_task.MarkTask(JsonStorage, id, entities.InProgress)
			if err != nil {
				return err
			}
		case "done":
			err := mark_task.MarkTask(JsonStorage, id, entities.Done)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(markCmd)
}
