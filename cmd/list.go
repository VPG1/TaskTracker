package cmd

import (
	"TaskTracker/internal/entities"
	"TaskTracker/internal/usecases/get_all_tasks"
	"TaskTracker/internal/usecases/get_tasks_by_status"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func printLine() {
	for i := 1; i < 70; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func printTasks(tasks []*entities.Task) {
	if len(tasks) == 0 {
		return
	}

	for _, task := range tasks {
		printLine()
		fmt.Printf("Task ID: %v\n", task.Id)
		fmt.Printf("Task Name: %s\n", task.Name)
		fmt.Printf("Task Description: %s\n", task.Description)
		fmt.Printf("Task Status: %s\n", task.Status)
		fmt.Printf("Created At: %s\n", task.CreatedAt.Format(time.RFC822))
		fmt.Printf("Updated At: %s\n", task.UpdatedAt.Format(time.RFC822))
	}
	printLine()
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [todo|in-progress|done]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Prints tasks to stdout.",
	Long: `This command prints tasks to stdout.
This command can print all task or only tasks with the selected status.
You can select status by entering an optional argument with task status.
If the command called without argument, all tasks are printed to stdout.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println("All tasks:")
			tasks, err := get_all_tasks.GetAllTasks(JsonStorage)
			if err != nil {
				return err
			}

			printTasks(tasks)
		} else if len(args) == 1 {
			switch args[0] {
			case "todo":
				tasks, err := get_tasks_by_status.GetTasksByStatus(JsonStorage, entities.NotDone)
				if err != nil {
					return err
				}

				printTasks(tasks)
			case "in-progress":
				tasks, err := get_tasks_by_status.GetTasksByStatus(JsonStorage, entities.InProgress)
				if err != nil {
					return err
				}

				printTasks(tasks)
			case "done":
				tasks, err := get_tasks_by_status.GetTasksByStatus(JsonStorage, entities.Done)
				if err != nil {
					return err
				}

				printTasks(tasks)
			default:
				return fmt.Errorf("invalid argument (avalible arguments [todo|in-progress|done]))")
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
